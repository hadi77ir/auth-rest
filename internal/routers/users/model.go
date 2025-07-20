package users

import (
	"auth-rest/internal/modules/users"
	"time"
)

type UserResponse struct {
	ID        int    `json:"id" example:"1"`
	Phone     string `json:"phone" example:"09123456789"`
	Role      string `json:"role" example:"admin"`
	Name      string `json:"name" example:"Admin"`
	CreatedAt string `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	UpdatedAt string `json:"updated_at" example:"2020-09-20T14:00:00+09:00"`
}

type UserUpdateRequest struct {
	Role string `json:"role" example:"admin"`
	Name string `json:"name" example:"Admin"`
}
type SelfUpdateRequest struct {
	Name string `json:"name" example:"Admin"`
}

type UserCreateRequest struct {
	Phone string `json:"phone" example:"09123456789"`
	Role  string `json:"role" example:"admin"`
	Name  string `json:"name" example:"Admin"`
}
type UserCreateResponse struct {
	ID      int  `json:"id" example:"1"`
	Success bool `json:"success" example:"true"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

func ToResponse(user users.UserModel) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Phone:     user.Phone,
		Role:      user.Role,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
