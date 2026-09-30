package sql

import (
	"errors"

	"github.com/jackc/pgx/v5"
	myerrors "go.lukasgrimm.me/homepage/internal/errors"
)

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return myerrors.NotFoundError(err)
	}

	return myerrors.InternalServerError(err)
}
