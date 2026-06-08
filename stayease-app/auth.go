package stayease

type Role string

const (
	Owner        Role = "Owner"
	Manager      Role = "Manager"
	Receptionist Role = "Receptionist"
	Cleaner      Role = "Cleaner"
)

var UsersDB = map[string]User{
	"owner":        {Username: "owner", Password: "ownerpass123", Role: Owner},
	"manager":      {Username: "manager", Password: "managerpass123", Role: Manager},
	"receptionist": {Username: "receptionist", Password: "receppass123", Role: Receptionist},
	"cleaner":      {Username: "cleaner", Password: "cleanerpass123", Role: Cleaner},
}

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
