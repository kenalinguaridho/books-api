package middleware

import (
	"github.com/kenalinguaridho/books-api/config"
	"github.com/kenalinguaridho/books-api/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil && err == http.ErrNoCookie{
			helper.Response(w, http.StatusUnauthorized, "Error user unauthorized", nil)
			return
		}

		// mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token)(interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				helper.Response(w, http.StatusUnauthorized, "Error user unauthorized", nil)
				return
			
			case jwt.ValidationErrorExpired:
				// token expired
				helper.Response(w, http.StatusUnauthorized, "Error token expired", nil)
				return
			default:
				helper.Response(w, http.StatusUnauthorized, "Error user unauthorized", nil)
				return
			}
		}

		if !token.Valid {
			helper.Response(w, http.StatusUnauthorized, "Error user unauthorized", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}