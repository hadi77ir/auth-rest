package run

import (
	"auth-rest/internal/app"
	"auth-rest/internal/log"
	"auth-rest/pkg/cmd"
	"context"
	"github.com/hadi77ir/go-logging"
	"github.com/urfave/cli/v3"
	"os"
	"os/signal"
	"syscall"
)

var RunCmd = &cli.Command{
	Name:        "run",
	Usage:       "run the system",
	Description: `this command runs the server.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
		},
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		globals := ctx.Value("globals").(*app.AppGlobals)
		logger := log.FromGlobals(globals)

		// handle sigint and sigterm
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithCancel(ctx)
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			cancelFunc()
		}()

		// run server
		err := cmd.Run(ctx, globals, cmd.RunArgs{
			ConfigPath: command.String("config"),
		})
		if err != nil {
			logger.Log(logging.FatalLevel, err)
		}
		return nil
	},
}
