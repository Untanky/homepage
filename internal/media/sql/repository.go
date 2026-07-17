package sql

import (
	"bytes"
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/internal/media"
)

type MediaRepository struct {
	db *pgx.Conn
}

func NewMediaRepository(db *pgx.Conn) *MediaRepository {
	return new(MediaRepository{
		db: db,
	})
}

func (repo *MediaRepository) GetAssetVersion(ctx context.Context, path string, scale float64, mimetype string) (media.AssetVersion, error) {
	row := repo.db.QueryRow(ctx, "SELECT av.data FROM media.asset_versions av JOIN media.assets a ON av.asset_id = a.id WHERE a.name = $1 AND av.scale = $2 AND av.mimetype = $3 LIMIT 1", path, scale, strings.ToUpper(mimetype))

	version := media.AssetVersion{
		Scale:     scale,
		MediaType: mimetype,
	}

	data := []byte{}

	if err := row.Scan(&data); err != nil {
		return media.AssetVersion{}, err
	}

	version.Buffer = bytes.NewBuffer(data)

	return version, nil
}

func (repo *MediaRepository) Create(ctx context.Context, asset media.Asset) error {
	batch := new(pgx.Batch)

	batch.Queue("INSERT INTO media.assets (id, name, width, height) VALUES ($1, $2, $3, $4)", asset.ID, asset.Name, asset.Width, asset.Height)

	for _, version := range asset.Versions {
		batch.Queue("INSERT INTO media.asset_versions (asset_id, scale, mimetype, data) VALUES ($1, $2, $3, $4)",
			asset.ID,
			version.Scale,
			version.MediaType,
			version.Buffer.Bytes(),
		)
	}

	results := repo.db.SendBatch(ctx, batch)
	defer results.Close()

	// Iterate to check errors
	for range asset.Versions {
		_, err := results.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
