package entity

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Permissions []Permission
}

func NewUser(email, first_name, last_name, password string, permissions []Permission) User {
	return User{
		Email:       email,
		FirstName:   first_name,
		LastName:    last_name,
		Password:    password,
		Permissions: permissions,
	}
}
