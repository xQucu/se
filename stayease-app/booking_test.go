package stayease

import "testing"

func TestCalculateBill(t *testing.T) {
	bill := CalculateBill(120.0, 4)
	if bill != 480.0 {
		t.Errorf("expected bill to be 480.00, got: %f", bill)
	}

	negativeBill := CalculateBill(120.0, -2)
	if negativeBill != 0 {
		t.Errorf("expected bill for negative duration to be 0, got %f", negativeBill)
	}
}
