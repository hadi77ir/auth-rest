package sms

import (
	"auth-rest/internal/app"
	"auth-rest/internal/log"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/hadi77ir/go-logging"
)

type SMSProvider interface {
	SendSMS(phoneNumber string, content string) error
}

type LoggingSMSProvider struct {
	logger logging.Logger
}

func (p *LoggingSMSProvider) SendSMS(phoneNumber string, content string) error {
	p.logger.Log(logging.InfoLevel, "Sending SMS: phoneNumber = ", phoneNumber, ", content = ", content)
	return nil
}

var _ SMSProvider = &LoggingSMSProvider{}

func NewLoggingProvider(globals *app.AppGlobals) (SMSProvider, error) {
	return &LoggingSMSProvider{logger: log.FromGlobals(globals)}, nil
}

func NewProviderByName(globals *app.AppGlobals, name string) (SMSProvider, error) {
	if name == "log" {
		return NewLoggingProvider(globals)
	}
	return nil, errors.New("invalid sms provider. only log is supported")
}

const SMSProviderKey = "smsprovider"

func Setup(globals *app.AppGlobals, providerName string) error {
	provider, err := NewProviderByName(globals, providerName)
	if err != nil {
		return err
	}
	globals.Set(SMSProviderKey, provider)
	return nil
}

func FromGlobals(globals *app.AppGlobals) SMSProvider {
	return app.Value[SMSProvider](globals, SMSProviderKey)
}

func FromContext(ctx fiber.Ctx) SMSProvider {
	return FromGlobals(app.FromContext(ctx))
}
