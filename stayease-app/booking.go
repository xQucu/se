package stayease

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
