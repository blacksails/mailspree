package mailspree

// AuthService is a common interface to token authentication. Authenticate
// takes a User and a password, and returns a token if they match. Validate
// takes a token and returns a user if it is valid
type AuthService interface {
	Authenticate(User, string) (string, error)
	Validate(string, UserService) (User, error)
}
