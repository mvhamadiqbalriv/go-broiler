package middleware

import (
	"context"
	"errors"
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/exception"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

func AuthenticateMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Retrieve the token from the Authorization header
		tokenString := getTokenFromHeader(r)
		if tokenString == "" {
			//return 
			fmt.Println("Token not found")
			panic(exception.NewUnauthorizedError("Token not found"))
		}

		// Verify the token
		token, err := verifyToken(tokenString)
		if err != nil {
			//return 
			fmt.Println("Token invalid")
			panic(exception.NewUnauthorizedError("Token invalid"))
		}

		// Extract user data from token claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Failed to extract token claims", http.StatusUnauthorized)
            return
        }

        // Access user data
        loggedUserId := claims["iss"].(string) // Assuming userID is stored in token claims

        // Set user data in context for downstream handlers to access
        ctx := r.Context()
        ctx = context.WithValue(ctx, "loggedUserId", loggedUserId)
        r = r.WithContext(ctx)


		// Call the next handler
		next(w, r, ps)
	}
}

func getTokenFromHeader(r *http.Request) string {
	// Retrieve the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Split the header value
	// Format: Bearer <token>
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	// Return the token
	return tokenParts[1]
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
