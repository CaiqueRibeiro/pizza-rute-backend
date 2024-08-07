package middlewares

import (
	"context"
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/configs"
)

type contextValues struct {
	Key   interface{}
	Value interface{}
}

// injects values in http.Request context, then returns an http.Handler with inserted values
func withValuesMiddleware(values []contextValues) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			for _, v := range values {
				r = r.WithContext(context.WithValue(r.Context(), v.Key, v.Value))
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Gets value to be injected and inject it in context trough "withValuesMiddleware"
func injectValues() func(next http.Handler) http.Handler {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	return withValuesMiddleware([]contextValues{
		{
			Key:   "jwt",
			Value: cfg.JWTSecret,
		},
		{
			Key:   "jwtExpiresIn",
			Value: cfg.JWTExpiresIn,
		},
	})
}

// Injects value and use it's as middleware in the main handler
func WithContext(handler http.HandlerFunc) http.Handler {
	withValuesInjected := injectValues()
	return withValuesInjected(handler)
}
