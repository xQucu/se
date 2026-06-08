package stayease

func HasPermission(role Role, action string) bool {
	switch role {
	case Owner:
		return true // Owner has access to everything
	case Manager:
		return action != "manage_users" // Manager has second most permissions
	case Receptionist:
		return action == "view_rooms" || action == "calculate_bill" || action == "manage_reservations" || action == "checkout_guest"
	case Cleaner:
		return action == "view_rooms" || action == "update_room_status"
	default:
		return false
	}
}
