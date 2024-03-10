package main

import (
	"log/slog"
	"os"
	"sso_pet/internal/app"
	"sso_pet/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func main() {
	cfg, err := config.ConfigLoad()

	if err != nil {
		panic(err)
	}

	log := setupLogger(cfg.Env)

	log.Info("application starting", slog.String("env", cfg.Env), slog.Int("port", cfg.GRPC.Port))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	runChan := make(chan error)

	go application.GRPCSrv.Run(runChan) 
	if err := <-runChan; err != nil {
		panic(err)
	}

	//TODO: launch gRPC-server
}
