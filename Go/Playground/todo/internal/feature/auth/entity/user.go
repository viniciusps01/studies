package entity

type User struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	Permissions []Permission `json:"permissions"`
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
