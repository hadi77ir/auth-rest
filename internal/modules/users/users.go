package users

import (
	"auth-rest/internal/dal"
	"auth-rest/internal/db/schema"
	"time"
)

type UserModel struct {
	ID        int
	Phone     string
	Verified  bool
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ListUsers(usersRepo dal.UsersRepository) ([]UserModel, error) {
	users, err := usersRepo.All()
	if err != nil {
		return nil, err
	}
	result := make([]UserModel, len(users))
	for i, user := range users {
		result[i] = ToModel(user)
	}
	return result, nil
}

func ToModel(user *schema.User) UserModel {
	return UserModel{
		ID:        int(user.ID),
		Phone:     user.Phone,
		Verified:  user.Verified,
		Role:      user.Role,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserByID(usersRepo dal.UsersRepository, userID int) (UserModel, error) {
	user, err := usersRepo.ById(uint(userID))
	if err != nil {
		return UserModel{}, err
	}
	return ToModel(user), nil
}

func UserByPhone(usersRepo dal.UsersRepository, phone string) (UserModel, error) {
	user, err := usersRepo.ByPhoneNumber(phone)
	if err != nil {
		return UserModel{}, err
	}
	return ToModel(user), nil
}

func CreateUser(userRepo dal.UsersRepository, user UserModel) (UserModel, error) {
	userSchema := &schema.User{Phone: user.Phone, Verified: user.Verified, Role: user.Role}
	err := userRepo.Create(userSchema)
	if err != nil {
		return UserModel{}, err
	}
	return ToModel(userSchema), nil
}

func UpdateUser(userRepo dal.UsersRepository, user UserModel) error {
	userSchema := &schema.User{Phone: user.Phone, Verified: user.Verified, Role: user.Role}
	userSchema.ID = uint(user.ID)
	return userRepo.Update(userSchema)
}

func DeleteUser(userRepo dal.UsersRepository, userID int) error {
	return userRepo.Delete(uint(userID))
}
