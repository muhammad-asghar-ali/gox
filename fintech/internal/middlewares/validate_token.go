package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/muhammad-asghar-ali/go/fintech/internal/types"
)

func VerifyAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			resp := types.ErrResponse{
				Code:    http.StatusUnauthorized,
				Message: "No authorization header",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(resp)
			return
		}

		token_string := strings.TrimPrefix(authorization, "Bearer ")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (any, error) {
			return []byte("TokenPassword"), nil
		})

		if err != nil || !token.Valid {
			resp := types.ErrResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid or expired token",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(resp)
			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			resp := types.ErrResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token data",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(resp)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
