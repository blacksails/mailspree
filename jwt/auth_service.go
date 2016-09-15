package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/blacksails/mailspree"
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthService is a mailspree.AuthService based on jwt.
type AuthService struct {
	PrivateKey string
}

// Authenticate takes a user and password, if they match, a jwt is returned. If
// not an error is returned.
func (as AuthService) Authenticate(u mailspree.User, password string) (string, error) {
	passwordOK := u.CheckPassword(password)
	if !passwordOK {
		return "", errors.New("Invalid username or password")
	}
	return as.generateToken(u.Username), nil
}

// Validate takes a jwt token string and validates it. If it passes validation,
// user is fetched and returned from the user service
func (as AuthService) Validate(tokenStr string, us mailspree.UserService) (mailspree.User, error) {
	var claims jwt.StandardClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.PrivateKey), nil
	})
	if err != nil || !token.Valid {
		return mailspree.User{}, errors.New("Invalid token")
	}
	u, err := us.Find(claims.Subject)
	if err != nil {
		return mailspree.User{}, errors.New("Valid token but user could not be found")
	}
	return u, nil
}

func (as AuthService) generateToken(username string) string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(14*24) * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(as.PrivateKey))
	if err != nil {
		log.Fatalf("There was an error generating a token: %v", err)
	}
	return tokenString
}
