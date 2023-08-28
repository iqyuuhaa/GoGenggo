package middlewares

import (
	"net/http"

	"gogenggo/utils"
)

func WrapHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-MD5, Content-Length, Host, User-Agent, Accept-Encoding, X-CSRF-Token, Authorization, Unix-Time, X-Signature")

		if r.Method == http.MethodOptions {
			return
		}

		ctx := utils.SetProccessTimeCtx(r.Context())
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
