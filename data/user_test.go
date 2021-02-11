package data

import "testing"

func TestCheckValidation(t *testing.T) {
	u := &User{
		Username:    "JeremiS",
		Email:       "jeremi@gmail.com",
		DateOfBirth: "",
	}

	err := u.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
