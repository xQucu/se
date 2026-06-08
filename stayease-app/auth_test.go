package stayease

import "testing"

func TestValidatePassword(t *testing.T) {
	if ValidatePassword("short") {
		t.Errorf("expected validation to fail for short password")
	}
	if !ValidatePassword("longpassword123") {
		t.Errorf("expected validation to pass for long password")
	}
}
