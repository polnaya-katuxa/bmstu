package openapi

import (
	"errors"
	"fmt"
	"net/http"

	myerrors "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
)

func toErrorResponse(err error, defaultMessage string) (ImplResponse, error) {
	var myErr *myerrors.InsufficientBalanceError

	switch {
	case errors.Is(err, myerrors.ErrPermissionDenied):
		return ImplResponse{
			Code: http.StatusInternalServerError,
			Body: ErrorResponse{
				Message:       "Database permission denied - slave write operation.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, ErrParams):
		return ImplResponse{
			Code: http.StatusBadRequest,
			Body: ErrorResponse{
				Message:       "Wrong request parameters.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrGet):
		return ImplResponse{
			Code: http.StatusUnauthorized,
			Body: ErrorResponse{
				Message:       "Authorization failed.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPaid):
		return ImplResponse{
			Code: http.StatusPaymentRequired,
			Body: ErrorResponse{
				Message:       "Post is available only after subscription.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrUncomment):
		return ImplResponse{
			Code: http.StatusForbidden,
			Body: ErrorResponse{
				Message:       "Operation can be proceeded only for your comments or comments under your posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrAuthor):
		return ImplResponse{
			Code: http.StatusForbidden,
			Body: ErrorResponse{
				Message:       "Operation can be proceeded only for authored posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrReact):
		return ImplResponse{
			Code: http.StatusForbidden,
			Body: ErrorResponse{
				Message:       "Operation can be proceeded only for not authored posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrLogin):
		return ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: ErrorResponse{
				Message:       "Incorrect login length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPassword):
		return ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: ErrorResponse{
				Message:       "Incorrect password length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrEmail):
		return ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: ErrorResponse{
				Message:       "Incorrect email format. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrExists):
		return ImplResponse{
			Code: http.StatusConflict,
			Body: ErrorResponse{
				Message:       "User already exists.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPerms):
		return ImplResponse{
			Code: http.StatusForbidden,
			Body: ErrorResponse{
				Message:       "Operation can be proceeded only by an admin.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrNotFound):
		return ImplResponse{
			Code: http.StatusNotFound,
			Body: ErrorResponse{
				Message:       "Not found.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrAutoDelete):
		return ImplResponse{
			Code: http.StatusBadRequest,
			Body: ErrorResponse{
				Message:       "Cannot delete yourself.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.As(err, &myErr):
		return ImplResponse{
			Code: http.StatusPaymentRequired,
			Body: ErrorResponse{
				Message:       fmt.Sprintf("Insufficient balance: needed %d, got %d.", myErr.Want, myErr.Got),
				SystemMessage: err.Error(),
			},
		}, nil
	default:
		return ImplResponse{
			Code: http.StatusInternalServerError,
			Body: ErrorResponse{
				Message:       defaultMessage,
				SystemMessage: err.Error(),
			},
		}, nil
	}
}
