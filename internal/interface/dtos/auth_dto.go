package dtos

type LoginUserDTO struct {
	StudentID string `json:"student_id"` // user's id
	Password  string `json:"password"`   // user's password
}

type LoginResponseDTO struct {
	AccessToken string `json:"access_token"`
}
