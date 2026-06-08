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

func AuthenticateUser(username, password string) (*User, bool) {
	u, exists := UsersDB[username]
	if !exists || u.Password != password {
		return nil, false
	}
	return &u, true
}
