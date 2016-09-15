package mailspree_test

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/blacksails/mailspree"
)

func TestEmailString(t *testing.T) {
	tests := []struct {
		in  mailspree.Email
		out string
	}{
		{mailspree.Email{Name: "Test", Address: "test@test.com"}, "Test <test@test.com>"},
		{mailspree.Email{Address: "test@test.com"}, "test@test.com"},
	}
	for _, test := range tests {
		actual := test.in.String()
		if expected := test.out; actual != expected {
			t.Errorf("Expected '%v' to equal '%v'", actual, expected)
		}
	}
}

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		in  mailspree.Email
		out bool
	}{
		{mailspree.Email{Address: "test@test.com"}, true},
		{mailspree.Email{Address: ""}, false},
		{mailspree.Email{Address: "test"}, false},
		{mailspree.Email{Address: "test@"}, false},
		{mailspree.Email{Address: "@test"}, false},
		{mailspree.Email{Address: "test@test"}, false},
	}
	for _, test := range tests {
		actual, err := govalidator.ValidateStruct(test.in)
		fmt.Printf("%v: %v, %v\n", test.in.String(), actual, err)
		expected := test.out
		if actual != expected {
			t.Errorf("Validity of email '%v' is expected to be %v", test.in.String(), expected)
		}
	}
}
