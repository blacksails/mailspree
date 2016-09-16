package mailspree

import "golang.org/x/crypto/bcrypt"

// User represents a user of the mailspree system
type User struct {
	Username     string
	PasswordHash []byte
}

// NewUser creates a new user from a username and a password.
func NewUser(username, password string) User {
	u := User{Username: username}
	u.SetPassword(password)
	return u
}

// CheckPassword checks if the given password is correct
func (u User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(p))
	return err == nil
}

// SetPassword takes a cleartext password and sets the user password.
func (u *User) SetPassword(p string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	u.PasswordHash = hash
}
