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
