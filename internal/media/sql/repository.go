package sql

import (
	"bytes"
	"context"

	"github.com/jackc/pgx/v5"
	"go.lukasgrimm.me/homepage/internal/database"
	"go.lukasgrimm.me/homepage/internal/media"
)

type assetRow struct {
	ID     media.AssetID `db:"id"`
	Path   string        `db:"path"`
	Alt    string        `db:"alt"`
	Width  uint          `db:"width"`
	Height uint          `db:"height"`

	Versions []media.VersionMetadata `db:"versions"`
}

type versionRow struct {
	AssetID  media.AssetID `db:"asset_id"`
	Mimetype string        `db:"mimetype"`
	Width    uint          `db:"width"`
	Height   uint          `db:"height"`
	Data     []byte        `db:"data"`
}

type MediaRepository struct {
	db database.Client
}

func NewMediaRepository(db database.Client) *MediaRepository {
	return new(MediaRepository{
		db: db,
	})
}

func (repo *MediaRepository) ListAssets(ctx context.Context, filter string) ([]media.Asset, error) {
	const listAssetsSQL = `
		SELECT a.id, a.path, a.alt, a.width, a.height,
		       json_agg(json_build_object('Mimetype', av.mimetype, 'Width', av.width, 'Height', av.height)) as versions
		FROM media.assets a
		JOIN media.asset_versions av ON a.id = av.asset_id
		WHERE starts_with(a.path, $1)
		GROUP BY a.id
	`

	result, err := repo.db.Query(ctx, listAssetsSQL, filter)
	if err != nil {
		return nil, err
	}

	rows, err := pgx.CollectRows(result, pgx.RowToAddrOfStructByName[assetRow])
	if err != nil {
		return nil, err
	}

	assets := make([]media.Asset, len(rows))

	for idx, row := range rows {
		assets[idx] = media.Asset{
			ID:       row.ID,
			Path:     row.Path,
			Alt:      row.Alt,
			Width:    row.Width,
			Height:   row.Height,
			Versions: row.Versions,
		}
	}

	return assets, nil
}

func (repo *MediaRepository) GetAsset(ctx context.Context, assetID media.AssetID) (media.Asset, error) {
	const getAssetSQL = `
		SELECT a.id, a.path, a.alt, a.width, a.height,
		       json_agg(json_build_object('Mimetype', av.mimetype, 'Width', av.width, 'Height', av.height)) as versions
		FROM media.assets a
		JOIN media.asset_versions av ON a.id = av.asset_id
		WHERE a.id = $1
		GROUP BY a.id
	`

	result, err := repo.db.Query(ctx, getAssetSQL, assetID)
	if err != nil {
		return media.Asset{}, err
	}

	row, err := pgx.CollectExactlyOneRow(result, pgx.RowToAddrOfStructByName[assetRow])
	if err != nil {
		return media.Asset{}, err
	}

	return media.Asset{
		ID:       row.ID,
		Path:     row.Path,
		Alt:      row.Alt,
		Width:    row.Width,
		Height:   row.Height,
		Versions: row.Versions,
	}, nil
}

func (repo *MediaRepository) GetVersion(ctx context.Context, path string, mimetype string, width uint) (media.Version, error) {
	const getVersionSQL = `
		SELECT a.id as asset_id, av.mimetype, av.width, av.height, av.data
		FROM media.assets a
		JOIN media.asset_versions av ON a.id = av.asset_id
		WHERE a.path = $1 AND av.mimetype = $2 AND av.width = $3
	`

	result, err := repo.db.Query(ctx, getVersionSQL, path, mimetype, width)
	if err != nil {
		return media.Version{}, err
	}

	row, err := pgx.CollectExactlyOneRow(result, pgx.RowToAddrOfStructByName[versionRow])
	if err != nil {
		return media.Version{}, err
	}

	return media.Version{
		AssetID: row.AssetID,
		VersionMetadata: media.VersionMetadata{
			Mimetype: row.Mimetype,
			Width:    row.Width,
			Height:   row.Height,
		},
		Size:    len(row.Data),
		Content: bytes.NewReader(row.Data),
	}, nil
}

func (repo *MediaRepository) Create(ctx context.Context, asset media.Asset, versions []media.Version) error {
	const (
		insertAssetSql        = `INSERT INTO media.assets (id, path, alt, width, height) VALUES ($1, $2, $3, $4, $5)`
		insertAssetVersionSql = `INSERT INTO media.asset_versions (asset_id, mimetype, width, height, data) VALUES ($1, $2, $3, $4, $5)`
	)

	batch := new(pgx.Batch)

	batch.Queue(insertAssetSql, asset.ID, asset.Path, asset.Alt, asset.Width, asset.Height)

	for _, version := range versions {
		data := make([]byte, version.Size)
		_, err := version.Content.Read(data)
		if err != nil {
			return err
		}

		batch.Queue(insertAssetVersionSql, asset.ID, version.Mimetype, version.Width, version.Height, data)
	}

	results := repo.db.SendBatch(ctx, batch)

	for i := 0; i < batch.Len(); i++ {
		if _, err := results.Exec(); err != nil {
			return err
		}
	}

	return nil
}
