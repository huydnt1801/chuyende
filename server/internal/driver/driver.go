package driver

type Driver struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    string `json:"-"`
}
