package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) IsValid() bool {
	isEmailValid := r.Email != ""
	isPasswordValid := r.Password != ""

	return isEmailValid && isPasswordValid
}
