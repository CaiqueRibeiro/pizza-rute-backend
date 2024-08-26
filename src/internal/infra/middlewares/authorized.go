package middlewares

import (
	"context"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/pkg/utils"
)

type ContextClaimsKey string

func Authorized(handler http.HandlerFunc) http.Handler {
	withValuesInjected := injectValues()
	return withValuesInjected(authMiddleware(http.HandlerFunc(handler)))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := r.Context().Value("jwt").(string)
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[len("Bearer "):]
		claims, err := utils.VerifyJWT(tokenString, []byte(secretKey))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var key ContextClaimsKey = "props"
		ctx := context.WithValue(r.Context(), key, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
