package user

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Confirmed   bool   `json:"confirmed"`
	Password    string `json:"-"`
}

type UserUpdate struct {
	FullName  string
	Password  string
	Confirmed *bool
}

type UserParams struct {
	PhoneNumber string
	ID          int
}
