package dal

import (
	"auth-rest/internal/db/schema"
	"errors"
	"gorm.io/gorm"
)

type UsersRepository interface {
	ById(id uint) (*schema.User, error)
	ByPhoneNumber(phone string) (*schema.User, error)
	All() ([]*schema.User, error)
	Update(user *schema.User) error
	Create(user *schema.User) error
	Delete(id uint) error
}

func convertErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}
	return err
}
