package setup

import (
	"auth-rest/internal/app"
	"auth-rest/internal/log"
	"auth-rest/pkg/cmd"
	"context"
	"github.com/hadi77ir/go-logging"
	"github.com/urfave/cli/v3"
)

var SetupCmd = &cli.Command{
	Name:        "setup",
	Usage:       "sets up the system",
	Description: `setup auto-migrates the database and adds a super-admin user.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
		},
		&cli.StringFlag{
			Name:    "superadmin-phone",
			Aliases: []string{"p"},
		},
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		globals := ctx.Value("globals").(*app.AppGlobals)
		logger := log.FromGlobals(globals)

		err := cmd.FirstSetup(ctx, globals, cmd.FirstSetupArgs{
			SuperAdminPhone: command.String("superadmin-phone"),
			ConfigPath:      command.String("config"),
		})
		if err != nil {
			logger.Log(logging.FatalLevel, err)
			return err
		}
		return nil
	},
}
