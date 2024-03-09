package UserType

import "github.com/google/uuid"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Id          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Phone       string    `json:"phone"`
}

type Userdata struct {
	User User          `json:"user"`
	Data []Credentials `json:"data"`
}

type CredentialID struct {
	Id uuid.UUID
}
