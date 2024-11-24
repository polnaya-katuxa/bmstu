package v2

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *CustomResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	if statusCode == http.StatusNoContent {
		http.SetCookie(r, &http.Cookie{
			Name:     "user-token",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			MaxAge:   0,
			Domain:   "localhost",
			Path:     "/",
		})
	}
	r.ResponseWriter.WriteHeader(statusCode)
}

func Middleware(prl interfaces.IProfileLogic, logger *zap.SugaredLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		c, err := r.Cookie("user-token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, err := prl.AuthByToken(ctx, c.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx = mycontext.UserToContext(ctx, user)
		ctx = mycontext.LoggerToContext(ctx, logger)

		r = r.WithContext(ctx)

		crw := &CustomResponseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(crw, r)
	})
}
