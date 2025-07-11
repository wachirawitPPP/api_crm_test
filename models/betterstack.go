package models

import (
	"linecrmapi/structs"
	"log/slog"
	"os"

	slogbetterstack "github.com/samber/slog-betterstack"
)

var Token string = os.Getenv("BETTERSTRACK_TOKEN")
var Endpoint string = os.Getenv("BETTERSTRACK_ENDPOINT")

func AddBetterstackLog() {
	logger := slog.New(
		slogbetterstack.Option{
			Token:    Token,
			Endpoint: Endpoint,
		}.NewBetterstackHandler(),
	)

	// Logger is ready to be used
	logger.Info("Hello from Better Stack!")
}

func AddBetterstackLoginLog(payload structs.ObjLogCustomerLogin) {
	logger := slog.New(
		slogbetterstack.Option{
			Token:    Token,
			Endpoint: Endpoint,
		}.NewBetterstackHandler(),
	)

	// Logger is ready to be used
	logger.Info("CRM LINE LOGIN", slog.Any("payload", payload))
}
