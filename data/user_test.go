package data

import "testing"

func TestCheckValidation(t *testing.T) {
	u := &User{
		Username:    "JeremiS",
		Email:       "jeremi@gmail.com",
		DateOfBirth: "01/01/1999",
	}

	err := u.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
