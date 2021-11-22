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

func TestInvalidRelationshipType(t *testing.T) {
	u := &User{
		Username:    "JeremiS",
		Email:       "jeremi@gmail.com",
		DateOfBirth: "01/01/1999",
		Status:      "Away",
	}

	err := u.Validate()

	if !(err != nil && err.Error() == "Key: 'User.Status' Error:Field validation for 'Status' failed on the 'isStatusType' tag") {
		t.Errorf("A status type of value %s passed but StatusType need to be between %s and %s", u.Status, Online, InGame)
	}
}
