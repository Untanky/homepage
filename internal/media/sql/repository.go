package sql

import (
	"bytes"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media"
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
	Content  []byte        `db:"content"`
}

type MediaRepository struct {
	db database.Client
}

func NewMediaRepository(db database.Client) *MediaRepository {
	return new(MediaRepository{
		db: db,
	})
}

func (repo *MediaRepository) GetAsset(ctx context.Context, assetID media.AssetID) (media.Asset, error) {
	const getAssetSQL = `
		SELECT a.id, a.path, a.alt, a.width, a.height,
		       json_agg(json_build_object('Mimetype', av.mimetype, 'Width', av.width, 'Height', av.height)) as versions
		FROM media.assets a
		JOIN media.asset_versions av ON a.id = av.asset_id
		WHERE a.id = $1
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
		SELECT a.id as asset_id, av.mimetype, av.width, av.height, av.content
		FROM media.assets a
		JOIN media.asset_versions av ON a.id = av.asset_id
		WHERE a.path = $1 AND av.mimetype = $2 AND width = $3
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
		Size:    len(row.Content),
		Content: bytes.NewReader(row.Content),
	}, nil
}
