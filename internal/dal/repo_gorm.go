package dal

import (
	"auth-rest/internal/app"
	"gorm.io/gorm"
)

type GormRepositories struct {
	users *GormUsersRepository
}

func (g *GormRepositories) Users() UsersRepository {
	return g.users
}

var _ Repositories = &GormRepositories{}

func SetupWithGorm(globals *app.AppGlobals, dbc *gorm.DB) (Repositories, error) {
	repos := NewGormRepositories(dbc)
	globals.Set(RepositoriesKey, repos)
	return repos, nil
}

func NewGormRepositories(dbc *gorm.DB) Repositories {
	return &GormRepositories{
		users: &GormUsersRepository{
			db: dbc,
		},
	}
}
