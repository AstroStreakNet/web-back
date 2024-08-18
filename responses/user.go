package responses

type Login struct {
	Token string `json:"token"`
}

type GetDetails struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DisplayName string `json:"display_name"`
}

type UpdateDetails struct {
	Updates []UpdatedDetail `json:"updates"`
}

type UpdatedDetail struct {
	Detail   string `json:"updated_detail"`
	NewValue string `json:"new_value"`
	OldValue string `json:"old_value"`
}
