package app

import (
	sessionStorage "github.com/Gradient-and-Co/sigma-school/internal/adapter/auth/adapter/storage/redis"
	"github.com/Gradient-and-Co/sigma-school/internal/adapter/auth/jwt"
	authPort "github.com/Gradient-and-Co/sigma-school/internal/adapter/auth/port"
	v1 "github.com/Gradient-and-Co/sigma-school/internal/adapter/delivery/http/v1"
	"github.com/Gradient-and-Co/sigma-school/internal/adapter/payment/yoomoney"
	repository "github.com/Gradient-and-Co/sigma-school/internal/adapter/repository/postgres"
	storage "github.com/Gradient-and-Co/sigma-school/internal/adapter/storage/minio"
	"github.com/Gradient-and-Co/sigma-school/internal/app/config"
	"github.com/Gradient-and-Co/sigma-school/internal/app/server"
	"github.com/Gradient-and-Co/sigma-school/internal/core/port"
	"github.com/Gradient-and-Co/sigma-school/internal/core/service"
	"github.com/Gradient-and-Co/sigma-school/pkg/database/postgres"
	"github.com/Gradient-and-Co/sigma-school/pkg/database/redis"
	"github.com/Gradient-and-Co/sigma-school/pkg/logging"
	"github.com/Gradient-and-Co/sigma-school/pkg/minio"
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
			postgres.NewPostgresDB,
			redis.NewClient,
			minio.NewClient,
			// repositories
			fx.Annotate(
				repository.NewUserRepo,
				fx.As(new(port.IUserRepository)),
			),
			fx.Annotate(
				repository.NewCourseRepo,
				fx.As(new(port.ICourseRepository)),
			),
			fx.Annotate(
				repository.NewSchoolRepo,
				fx.As(new(port.ISchoolRepository)),
			),
			fx.Annotate(
				repository.NewLessonRepo,
				fx.As(new(port.ILessonRepository)),
			),
			fx.Annotate(
				repository.NewReviewRepo,
				fx.As(new(port.IReviewRepository)),
			),
			fx.Annotate(
				repository.NewStatRepo,
				fx.As(new(port.IStatRepository)),
			),
			fx.Annotate(
				storage.NewObjectStorage,
				fx.As(new(port.IObjectStorage)),
			),
			fx.Annotate(
				jwt.NewAuthProvider,
				fx.As(new(port.IAuthProvider)),
			),
			fx.Annotate(
				sessionStorage.NewSessionStorage,
				fx.As(new(authPort.ISessionStorage)),
			),
			fx.Annotate(
				yoomoney.NewPaymentGateway,
				fx.As(new(port.IPaymentGateway)),
			),
			// services
			fx.Annotate(
				service.NewUserService,
				fx.As(new(port.IUserService)),
			),
			fx.Annotate(
				service.NewCourseService,
				fx.As(new(port.ICourseService)),
			),
			fx.Annotate(
				service.NewSchoolService,
				fx.As(new(port.ISchoolService)),
			),
			fx.Annotate(
				service.NewLessonService,
				fx.As(new(port.ILessonService)),
			),
			fx.Annotate(
				service.NewReviewService,
				fx.As(new(port.IReviewService)),
			),
			fx.Annotate(
				service.NewMediaService,
				fx.As(new(port.IMediaService)),
			),
			fx.Annotate(
				service.NewStatService,
				fx.As(new(port.IStatService)),
			),
			fx.Annotate(
				service.NewPaymentService,
				fx.As(new(port.IPaymentService)),
			),
			fx.Annotate(
				service.NewAuthTokenService,
				fx.As(new(port.IAuthTokenService)),
			),
		),
		fx.Supply(cfg, &cfg.Redis, &cfg.Postgres, &cfg.JWT,
			&cfg.Minio, &cfg.Yoomoney, &cfg.Web, logger),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
