package v1

import (
	"errors"
	"fmt"
	"net/http"

	myerrors "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	openapi "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/openapi/v1"
)

func toErrorResponse(err error, defaultMessage string) (openapi.ImplResponse, error) {
	var myErr *myerrors.InsufficientBalanceError

	switch {
	case errors.Is(err, myerrors.ErrGet):
		return openapi.ImplResponse{
			Code: http.StatusUnauthorized,
			Body: openapi.ErrorResponse{
				Message:       "Authorization failed.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPaid):
		return openapi.ImplResponse{
			Code: http.StatusPaymentRequired,
			Body: openapi.ErrorResponse{
				Message:       "Post is available only after subscription.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrUncomment):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only for your comments or comments under your posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrAuthor):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only for authored posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrReact):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only for not authored posts.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrLogin):
		return openapi.ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: openapi.ErrorResponse{
				Message:       "Incorrect login length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPassword):
		return openapi.ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: openapi.ErrorResponse{
				Message:       "Incorrect password length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrEmail):
		return openapi.ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: openapi.ErrorResponse{
				Message:       "Incorrect email format. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrExists):
		return openapi.ImplResponse{
			Code: http.StatusConflict,
			Body: openapi.ErrorResponse{
				Message:       "User already exists.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrPerms):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only by an admin.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrNotFound):
		return openapi.ImplResponse{
			Code: http.StatusNotFound,
			Body: openapi.ErrorResponse{
				Message:       "Not found.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrAutoDelete):
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.ErrorResponse{
				Message:       "Cannot delete yourself.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.As(err, &myErr):
		return openapi.ImplResponse{
			Code: http.StatusPaymentRequired,
			Body: openapi.ErrorResponse{
				Message:       fmt.Sprintf("Insufficient balance: needed %d, got %d.", myErr.Want, myErr.Got),
				SystemMessage: err.Error(),
			},
		}, nil
	default:
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: openapi.ErrorResponse{
				Message:       defaultMessage,
				SystemMessage: err.Error(),
			},
		}, nil
	}
}
