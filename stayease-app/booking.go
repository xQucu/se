package stayease

import "errors"

var rolePermissions = map[Role]map[string]bool{
	Owner: {
		"manage_rooms":        true,
		"manage_users":        true,
		"view_rooms":          true,
		"update_room_status":  true,
		"calculate_bill":      true,
		"manage_reservations": true,
		"checkout_guest":      true,
	},
	Manager: {
		"manage_rooms":        true,
		"view_rooms":          true,
		"update_room_status":  true,
		"calculate_bill":      true,
		"manage_reservations": true,
		"checkout_guest":      true,
	},
	Receptionist: {
		"view_rooms":          true,
		"calculate_bill":      true,
		"manage_reservations": true,
		"checkout_guest":      true,
	},
	Cleaner: {
		"view_rooms":         true,
		"update_room_status": true,
	},
}

func HasPermission(role Role, action string) bool {
	perms, exists := rolePermissions[role]
	if !exists {
		return false
	}
	return perms[action]
}

func CreateRoom(role Role, roomNo string) error {
	if !HasPermission(role, "manage_rooms") {
		return errors.New("unauthorized to manage rooms")
	}
	return nil
}

type Room struct {
	ID     string
	Number string
	Status string
	Rate   float64
}

func CheckoutRoom(role Role, r *Room) (*Room, error) {
	if !HasPermission(role, "checkout_guest") {
		return nil, errors.New("unauthorized to checkout guest")
	}
	r.Status = "Needs Cleaning"
	return r, nil
}
