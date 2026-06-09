package stayease

import "testing"

func TestCalculateBill(t *testing.T) {
	bill := CalculateBill(120.0, 4)
	if bill != 480.0 {
		t.Errorf("expected bill to be 480.00, got: %f", bill)
	}
}
