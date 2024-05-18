package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	email := "a"

	if err := ValidateEmail(email); err == nil {
		t.Errorf("email: %q, should be invalid", email)
	}

	email = "valid@email.com"

	if err := ValidateEmail(email); err != nil {
		t.Errorf("email: %q, should be valid", email)
	}
}
