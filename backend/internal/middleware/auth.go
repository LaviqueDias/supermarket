package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/LaviqueDias/supermarket/pkg/utils"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UserTypeKey contextKey = "user_type"
	UserRoleKey contextKey = "user_role"
)

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

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserTypeKey, claims.UserType)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func EmployeeOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userType, ok := r.Context().Value(UserTypeKey).(string)
		if !ok || userType != "employee" {
			utils.RespondError(w, http.StatusForbidden, "Acesso negado. Apenas funcionários podem acessar este recurso.")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ClientOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userType, ok := r.Context().Value(UserTypeKey).(string)
		if !ok || userType != "client" {
			utils.RespondError(w, http.StatusForbidden, "Acesso negado. Apenas clientes podem acessar este recurso.")
			return
		}
		next.ServeHTTP(w, r)
	})
}