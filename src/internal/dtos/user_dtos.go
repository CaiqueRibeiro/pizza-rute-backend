package dtos

type CreateUserInput struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PhotoUrl     string `json:"photo_url,omitempty"`
	JobPosition  string `json:"job_position"`
	TempPassword string `json:"temp_password"`
}
