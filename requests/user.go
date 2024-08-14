package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
