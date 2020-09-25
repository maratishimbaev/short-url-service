package middleware

import (
	"github.com/kataras/golog"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			requestNumber, ok := r.Context().Value("id").(string)

			err := recover()
			if err != nil {
				if ok {
					golog.Errorf("#%s: panic: %s", requestNumber, err.(error).Error())
				} else {
					golog.Errorf("panic: %s", err.(error).Error())
				}

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
