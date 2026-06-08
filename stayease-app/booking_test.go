package stayease

import "testing"

func TestHasPermission(t *testing.T) {
	if !HasPermission(Cleaner, "update_room_status") {
		t.Errorf("cleaner should be allowed to update room status")
	}
	if HasPermission(Cleaner, "calculate_bill") {
		t.Errorf("cleaner should not be allowed to calculate bill")
	}
	if !HasPermission(Receptionist, "calculate_bill") {
		t.Errorf("receptionist should be allowed to calculate bill")
	}
	if HasPermission(Receptionist, "manage_rooms") {
		t.Errorf("receptionist should not be allowed to manage rooms")
	}
}

func TestCreateRoom(t *testing.T) {
	if !CreateRoom(Owner, "104") {
		t.Errorf("owner should be able to create a room")
	}
	if CreateRoom(Receptionist, "105") {
		t.Errorf("receptionist should not be able to create a room")
	}
}
