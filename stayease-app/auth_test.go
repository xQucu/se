package stayease

import "testing"

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"short", false},
		{"1234567", false},
		{"12345678", true},
		{"verylongpassword", true},
	}
	for _, tt := range tests {
		if result := ValidatePassword(tt.password); result != tt.expected {
			t.Errorf("ValidatePassword(%q) = %v; expected %v", tt.password, result, tt.expected)
		}
	}
}

func TestAuthenticateUser(t *testing.T) {
	u, ok := AuthenticateUser("owner", "ownerpass123")
	if !ok || u == nil {
		t.Errorf("expected authentication to succeed")
	}
	if u.Role != Owner {
		t.Errorf("expected role to be Owner, got %v", u.Role)
	}

	_, badOk := AuthenticateUser("owner", "wrongpassword")
	if badOk {
		t.Errorf("expected authentication to fail for wrong password")
	}
}
