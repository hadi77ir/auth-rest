package cmd

import (
	"auth-rest/cmd/run"
	"auth-rest/cmd/setup"
	"auth-rest/internal/app"
	"auth-rest/internal/log"
	"context"
	"github.com/hadi77ir/go-logging"
	"github.com/urfave/cli/v3"
	"go.uber.org/automaxprocs/maxprocs"
)

func BeforeCommand(ctx context.Context, c *cli.Command) (context.Context, error) {
	globals := app.NewGlobals()
	logger := log.Setup(globals)

	_, err := maxprocs.Set()
	if err != nil {
		logger.Log(logging.ErrorLevel, "error setting MAXPROCS: ", err)
	}
	return context.WithValue(ctx, "globals", globals), nil
}

var RootCmd = &cli.Command{
	Name:   "auth-rest",
	Before: BeforeCommand,
}

func init() {
	// add subcommands
	RootCmd.Commands = []*cli.Command{run.RunCmd, setup.SetupCmd}
}
