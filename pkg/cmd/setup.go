package cmd

import (
	"auth-rest/internal/app"
	"auth-rest/internal/db"
	"auth-rest/internal/db/schema"
	"auth-rest/internal/modules/auth"
	"auth-rest/internal/utils"
	"context"
	"errors"
)

type FirstSetupArgs struct {
	SuperAdminPhone string
	ConfigPath      string
}

func FirstSetup(context context.Context, globals *app.AppGlobals, setupArgs FirstSetupArgs) error {
	if setupArgs.SuperAdminPhone == "" {
		return errors.New("super admin phone is required")
	}
	if !utils.IsValidPhone(setupArgs.SuperAdminPhone) {
		return errors.New("super admin phone is not valid")
	}
	_, _, dbc, repos, err := CommonInit(context, globals, setupArgs.ConfigPath)
	if err != nil {
		return err
	}
	err = db.AutoMigrate(dbc)
	if err != nil {
		return err
	}

	err = repos.Users().Create(&schema.User{
		Phone:    setupArgs.SuperAdminPhone,
		Role:     auth.RoleSuper,
		Verified: true,
	})
	if err != nil {
		return err
	}

	return nil
}
