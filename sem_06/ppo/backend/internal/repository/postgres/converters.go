package postgres

import (
	"errors"

	myerrors "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func toErrorResponse(err error) error {
	var pgErr *pgconn.PgError
	switch {
	case errors.As(err, &pgErr) && pgErr.Code == "25006":
		return myerrors.ErrPermissionDenied
	default:
		return err
	}
}
