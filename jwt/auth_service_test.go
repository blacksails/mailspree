package jwt_test

import (
	"testing"

	"github.com/blacksails/mailspree"
	"github.com/blacksails/mailspree/jwt"
)

func TestAuthServiceAuthenticate(t *testing.T) {
	as := jwt.AuthService{PrivateKey: "test"}
	pw := "secret"
	u := mailspree.NewUser("test", pw)
	_, err := as.Authenticate(u, pw)
	if err != nil {
		t.Error("should be able to authenticate with correct pw")
	}
	_, err = as.Authenticate(u, "wrongpw")
	if err == nil {
		t.Error("should not be able to authenticate with wrong pw")
	}
}

func TestAuthServiceValidate(t *testing.T) {
	as := jwt.AuthService{PrivateKey: "test"}
	pw := "secret"
	u := mailspree.NewUser("test", pw)
	token, _ := as.Authenticate(u, pw)
	us := mailspree.SimpleUserService{User: u}

	vu, err := as.Validate(token, us)
	if vu.Username != u.Username {
		t.Error("validate should return the correct user from the token")
	}
	if err != nil {
		t.Error("should not return error on successful validation")
	}

	vu, err = as.Validate("invalidtoken", us)
	if err == nil {
		t.Error("should return error if validating invalid token")
	}

	vu, err = as.Validate(token, mailspree.SimpleUserService{})
	if err == nil {
		t.Error("should return error if the user from the token couldn't be found")
	}
}
