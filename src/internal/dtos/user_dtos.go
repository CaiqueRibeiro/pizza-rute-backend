package dtos

type CreateUserInput struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PhotoUrl     string `json:"photo_url,omitempty"`
	JobPosition  string `json:"job_position"`
	TempPassword string `json:"temp_password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
}
