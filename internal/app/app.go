package app

import (
	v1 "github.com/Gradient-and-Co/sigma-school/internal/adapter/delivery/http/v1"
	"github.com/Gradient-and-Co/sigma-school/internal/app/config"
	"github.com/Gradient-and-Co/sigma-school/internal/app/server"
	"github.com/Gradient-and-Co/sigma-school/pkg/logging"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func RunWeb() {
	cfg := config.GetConfig()
	log.Println("config is loaded")

	logger, err := logging.NewLogger(&cfg.Logging)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.Info("logger initialized")
	logger.Info("application startup")

	fx.New(
		fx.Provide(
			server.NewServer,
			server.NewGinRouter,
			v1.NewHandler,
		),
		fx.Supply(cfg, &cfg.Web, logger),
		fx.Invoke(func(*http.Server) {}),
		fx.NopLogger,
	).Run()
}
