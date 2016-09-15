package mailspree

import "errors"

// SimpleUserService only has a single user. A lookup for any other user will
// return an error.
type SimpleUserService struct {
	User User
}

// Find checks if the given username matches the one of the User set in the
// SimpleUserService. If it does, the user is returned otherwise we get a not
// found error.
func (us SimpleUserService) Find(username string) (User, error) {
	if us.User.Username != username {
		return User{}, errors.New("No user was found")
	}
	return us.User, nil
}
