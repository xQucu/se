package stayease

import "testing"

func TestHasPermission(t *testing.T) {
	// Cleaner can view/update room status, cannot calculate bill or manage rooms
	if !HasPermission(Cleaner, "update_room_status") {
		t.Errorf("cleaner should be allowed to update room status")
	}
	if HasPermission(Cleaner, "calculate_bill") {
		t.Errorf("cleaner should not be allowed to calculate bill")
	}

	// Receptionist can calculate bill, check in, cannot manage rooms
	if !HasPermission(Receptionist, "calculate_bill") {
		t.Errorf("receptionist should be allowed to calculate bill")
	}
	if HasPermission(Receptionist, "manage_rooms") {
		t.Errorf("receptionist should not be allowed to manage rooms")
	}
}
