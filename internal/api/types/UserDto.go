package types

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}
