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

type UpdateDetails struct {
	Updates []Update `json:"updates"`
}

type Update struct {
	Field    string `json:"field"`
	NewValue string `json:"new_value"`
}
