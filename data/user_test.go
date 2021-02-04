package data

import "testing"

func TestCheckValidation(t *testing.T) {
	u := &User{
		Username: "sickboy",
		Email:    "jeremi@gmail.com",
	}

	err := u.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
