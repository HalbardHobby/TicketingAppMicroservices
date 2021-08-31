package routes

import (
	"net/http"
	"os"

	"github.com/HalbardHobby/TicketingAppMicroservices/auth/data"
	"github.com/golang-jwt/jwt"
)

func generateAuthenticationCookie(user data.User) (*http.Cookie, error) {
	// Generate JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}
	// Save on session
	return &http.Cookie{
		Name:  "session",
		Value: token,
		Path:  "/",
	}, nil

}
