package models

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterUserRequest) IsValid() bool {

	isUsernameValid := r.Username != "" && len(r.Username) <= 100
	isEmailValid := r.Email != ""
	isPasswordValid := r.Password != ""

	return isUsernameValid && isEmailValid && isPasswordValid
}
