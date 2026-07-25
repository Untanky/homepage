package sql

import (
	"bytes"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media"
)

type MediaRepository struct {
	db database.Client
}

func NewMediaRepository(db database.Client) *MediaRepository {
	return new(MediaRepository{
		db: db,
	})
}

func (repo *MediaRepository) GetAsset(ctx context.Context, id media.AssetID) (media.Asset, error) {
	const (
		queryAsset = `
			SELECT a.name, a.width, a.height
			FROM media.assets a
			WHERE a.id = $1
			LIMIT 1
		`
		queryAssetVersion = `
			SELECT av.scale, av.mimetype
			FROM media.asset_versions av
			WHERE av.asset_id = $1
		`
	)

	asset := media.Asset{}

	row := repo.db.QueryRow(ctx, queryAsset, id)

	if err := row.Scan(&asset.Name, &asset.Width, &asset.Height); err != nil {
		return media.Asset{}, err
	}

	result, err := repo.db.Query(ctx, queryAssetVersion, id)
	if err != nil {
		return media.Asset{}, err
	}

	for result.Next() {
		assetVersion := media.AssetVersion{}
		if err := result.Scan(&assetVersion.Scale, &assetVersion.MediaType); err != nil {
			return media.Asset{}, err
		}

		asset.Versions = append(asset.Versions, assetVersion)
	}

	return asset, nil
}

func (repo *MediaRepository) GetAssetVersion(ctx context.Context, path string, scale float64, mimetype string) (media.AssetVersion, error) {
	const queryAssetVersion = `
    SELECT av.data
    FROM media.asset_versions av
    JOIN media.assets a ON av.asset_id = a.id
    WHERE a.name = $1
      AND av.scale = $2
      AND av.mimetype = $3
    LIMIT 1
	`

	row := repo.db.QueryRow(ctx, queryAssetVersion, path, scale, mimetype)

	data := []byte{}

	if err := row.Scan(&data); err != nil {
		return media.AssetVersion{}, err
	}

	return media.AssetVersion{
		Scale:     scale,
		MediaType: mimetype,
		Buffer:    bytes.NewBuffer(data),
	}, nil
}

func (repo *MediaRepository) Create(ctx context.Context, asset media.Asset) error {
	const (
		insertAssetQuery = `
			INSERT INTO media.assets (id, name, width, height)
			VALUES ($1, $2, $3, $4)
		`
		insertAssetVersionQuery = `
			INSERT INTO media.asset_versions (asset_id, scale, mimetype, data)
			VALUES ($1, $2, $3, $4)
		`
	)

	batch := new(pgx.Batch)

	batch.Queue(insertAssetQuery, asset.ID, asset.Name, asset.Width, asset.Height)

	for _, version := range asset.Versions {
		batch.Queue(insertAssetVersionQuery, asset.ID, version.Scale, version.MediaType, version.Buffer.Bytes())
	}

	results := repo.db.SendBatch(ctx, batch)
	defer results.Close()

	// Iterate to check errors
	for i := 0; i < batch.Len(); i++ {
		if _, err := results.Exec(); err != nil {
			return err
		}
	}

	return nil
}
