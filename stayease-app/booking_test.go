package stayease

import "testing"

func TestHasPermission(t *testing.T) {
	if !HasPermission(Cleaner, "update_room_status") {
		t.Errorf("cleaner should be allowed to update room status")
	}
	if HasPermission(Cleaner, "calculate_bill") {
		t.Errorf("cleaner should not be allowed to calculate bill")
	}
}

func TestCheckoutRoom(t *testing.T) {
	r := Room{Status: "Occupied"}
	
	// Cleaner cannot checkout guest
	err := CheckoutRoom(Cleaner, &r)
	if err == nil {
		t.Errorf("cleaner should not be authorized to checkout guest")
	}

	// Receptionist can checkout guest
	err = CheckoutRoom(Receptionist, &r)
	if err != nil {
		t.Errorf("receptionist should be authorized to checkout guest, got: %v", err)
	}
	if r.Status != "Needs Cleaning" {
		t.Errorf("expected room status to be Needs Cleaning, got %s", r.Status)
	}
}
