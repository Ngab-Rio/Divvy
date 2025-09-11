package dto

type AuthLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthRegisterRequest struct {
	Username string `json:"username" validate:"required,min=4"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
}