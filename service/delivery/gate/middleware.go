package gate

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/smartm2m/models"
	"github.com/dgrijalva/jwt-go"
)

func Protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			StatusCode int
			Data       interface{}
			Message    string
		}{
			http.StatusUnauthorized,
			nil,
			"Unauthorized",
		}

		// Get the authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp.Message = "Should has Authorization header"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Check if the authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			resp.Message = "Authorization should be Bearer format"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Extract the token from the authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Here you would verify that the signing method is correct
			// and that the key used to sign the token is valid
			// For simplicity, we'll just return an arbitrary key here
			return []byte(utstring.GetEnv(models.AppApiSecret)), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			resp.Message = "Cannot get the cookies"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			resp.Message = "Cookies is not valid"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Get the userId claim from the token
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			resp.Message = "Cannot get the session"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Add the userId claim to the request context
		var session models.Session
		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			resp.Message = "Session should be numeric format"
			ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
		}
		session.UserID = userID
		ctx := context.WithValue(r.Context(), "session", session)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
