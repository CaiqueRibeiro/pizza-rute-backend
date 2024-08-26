package handlers

import (
	"net/http"

	"github.com/CaiqueRibeiro/pizza-rute/src/internal/infra/middlewares"
	"github.com/CaiqueRibeiro/pizza-rute/src/internal/services"
	"github.com/golang-jwt/jwt/v5"
)

func IsLoggedUserAllowed(r *http.Request) bool {
	var claimsKey middlewares.ContextClaimsKey = "props"
	claims := r.Context().Value(claimsKey).(jwt.MapClaims)
	return services.AdminPermission.IsAllowedToExecute(claims["role"].(string))
}
