package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*type DeleteUserDTO struct {
	UserId uuid.UUID `json:"id"`
}*/

/*type GetUserDTO struct {
	UserId uuid.UUID `json:"id"`
}*/
