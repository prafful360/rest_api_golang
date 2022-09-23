package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	HashPassword []byte `json:"hash_password"`
	Password     string `json:"password"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
