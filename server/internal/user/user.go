package user

type User struct {
	ID          int    `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Confirmed   bool   `json:"confirmed,omitempty"`
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
