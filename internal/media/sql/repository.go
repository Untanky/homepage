package sql

import (
	"context"

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
