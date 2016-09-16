package mailspree_test

import (
	"testing"

	"github.com/blacksails/mailspree"
)

func TestSimpleUserService(t *testing.T) {
	username := "test"
	user := mailspree.User{Username: username}
	us := mailspree.SimpleUserService{User: user}
	u, err := us.Find("test")
	if u.Username != user.Username {
		t.Error("we didn't find the right user")
	}
	if err != nil {
		t.Error("there should be no problem finding the user")
	}
	u, err = us.Find("notauser")
	if err == nil {
		t.Error("that user shouldn't be there")
	}
}
