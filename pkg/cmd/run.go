package cmd

import (
	"auth-rest/internal/app"
	"auth-rest/internal/modules/sms"
	"auth-rest/internal/routers/errors"
	"auth-rest/internal/setup"
	"auth-rest/internal/storage"
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/hadi77ir/go-logging"
	"net"
)

type RunArgs struct {
	ConfigPath string
}

func Run(context context.Context, globals *app.AppGlobals, runArgs RunArgs) error {
	cfg, logger, _, _, err := CommonInit(context, globals, runArgs.ConfigPath)
	if err != nil {
		return err
	}

	// setup sms provider
	err = sms.Setup(globals, cfg.SMSProvider)
	if err != nil {
		return err
	}
	logger.Log(logging.DebugLevel, "done setting up sms provider")

	_, err = storage.Setup(globals, nil)
	if err != nil {
		return err
	}

	server := fiber.New(fiber.Config{
		AppName:      "auth-rest",
		ErrorHandler: errors.HandleError,
	})

	// setup handlers (routes and middlewares)
	err = setup.SetupHandlers(globals, server)
	if err != nil {
		return err
	}
	logger.Log(logging.DebugLevel, "done setting up handlers")

	listenAddr := ":4000"
	if cfg.ListenAddr != "" {
		listenAddr = cfg.ListenAddr
	}
	err = server.Listen(listenAddr, fiber.ListenConfig{
		DisableStartupMessage: true,
		ListenerAddrFunc: func(addr net.Addr) {
			logger.Log(logging.InfoLevel, "listening on ", listenAddr)
		},
		OnShutdownSuccess: func() {
			logger.Log(logging.InfoLevel, "successfully shut down")
		},
		GracefulContext: context,
	})
	if err != nil {
		return err
	}
	return nil
}
