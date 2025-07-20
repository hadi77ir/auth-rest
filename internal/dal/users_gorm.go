package dal

import (
	"auth-rest/internal/db/schema"
	"gorm.io/gorm"
)

type GormUsersRepository struct {
	db *gorm.DB
}

func (g *GormUsersRepository) All() ([]*schema.User, error) {
	var users []*schema.User
	tx := g.db.Find(users)
	return users, convertErr(tx.Error)
}

func (g *GormUsersRepository) Create(user *schema.User) error {
	return g.db.Create(user).Error
}

func (g *GormUsersRepository) Update(user *schema.User) error {
	return g.db.Save(user).Error
}

func (g *GormUsersRepository) Delete(id uint) error {
	return convertErr(g.db.Delete(&schema.User{}, id).Error)
}

func (g *GormUsersRepository) ById(u uint) (*schema.User, error) {
	var user schema.User
	tx := g.db.Where("id = ?", u).Take(user)
	if tx.Error != nil {
		return nil, convertErr(tx.Error)
	}
	return &user, nil
}

func (g *GormUsersRepository) ByPhoneNumber(phone string) (*schema.User, error) {
	var user schema.User
	tx := g.db.Where("phone = ?", phone).Take(user)
	if tx.Error != nil {
		return nil, convertErr(tx.Error)
	}
	return &user, nil
}

var _ UsersRepository = &GormUsersRepository{}
