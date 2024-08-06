package entities

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhotoUrl    string `json:"photo_url"`
	JobPosition string `json:"job_position"`
}
