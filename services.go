package mailspree

// MailingProvider abstacts different mailing providers.
type MailingProvider interface {
	SendEmail(Message) error
}

// AuthService is a common interface to token authentication. Authenticate
// takes a User and a password, and returns a token if they match. Validate
// takes a token and returns a user if it is valid
type AuthService interface {
	Authenticate(User, string) (string, error)
	Validate(string, UserService) (User, error)
}

// UserService is a common interface for retrieving Users from usernames. This
// will make it easy to swap out with a database.
type UserService interface {
	Find(string) (User, error)
}
