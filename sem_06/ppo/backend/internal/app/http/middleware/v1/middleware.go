package v1

import (
	"net/http"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"go.uber.org/zap"
)

func Middleware(prl interfaces.IProfileLogic, logger *zap.SugaredLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := prl.AuthByToken(ctx, r.Header.Get("User-Token"))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx = mycontext.UserToContext(ctx, user)
		ctx = mycontext.LoggerToContext(ctx, logger)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
