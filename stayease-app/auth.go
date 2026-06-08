package stayease

type Role string

const (
	Owner        Role = "Owner"
	Manager      Role = "Manager"
	Receptionist Role = "Receptionist"
	Cleaner      Role = "Cleaner"
)

func ValidatePassword(password string) bool {
	return len(password) >= 8
}
