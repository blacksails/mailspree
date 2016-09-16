package mailspree_test

import (
	"testing"

	"github.com/blacksails/mailspree"

	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	username := "test"
	password := "somesecret"
	u := mailspree.NewUser(username, password)
	if u.Username != username {
		t.Error("input username should be set on the user")
	}
	if bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) != nil {
		t.Error("input password should be set to the bcrypt hash of the password")
	}
}

func TestUserCheckPassword(t *testing.T) {
	username := "test"
	password := "somesecret"
	u := mailspree.NewUser(username, password)
	if !u.CheckPassword(password) {
		t.Error("should return true when given correct password")
	}
	if u.CheckPassword("invalidpassword") {
		t.Error("should return false when give invalid password")
	}
}

func TestUserSetPassword(t *testing.T) {
	username := "test"
	password := "somesecret"
	u := mailspree.NewUser(username, password)
	newpwd := "supersecret"
	u.SetPassword(newpwd)
	if bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil {
		t.Error("the old password should not work anymore")
	}
	if bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(newpwd)) != nil {
		t.Error("the new password should work")
	}
}
