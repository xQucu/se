package stayease

type User struct {
	Username string
	Password string
	Role     Role
}

var UsersDB = map[string]User{
	"owner":        {Username: "owner", Password: "ownerpass123", Role: Owner},
	"manager":      {Username: "manager", Password: "managerpass123", Role: Manager},
	"receptionist": {Username: "receptionist", Password: "receppass123", Role: Receptionist},
	"cleaner":      {Username: "cleaner", Password: "cleanerpass123", Role: Cleaner},
}
