package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type contextKey string

const ClientIDKey contextKey = "client_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Token não fornecido")
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		ctx := context.WithValue(r.Context(), ClientIDKey, claims.ClientID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
