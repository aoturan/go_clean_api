package main

import (
	"context"
	"flag"
	"github.com/aoturan/go_clean_api/api"
	"github.com/aoturan/go_clean_api/pkg/session"
	"github.com/aoturan/go_clean_api/pkg/user"
	"github.com/aoturan/go_clean_api/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	envConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("cannot load config")
	}

	listenAddr := flag.String("listenAddr", envConfig.Port, "listen address of the http server")
	flag.Parse()

	client, cancel, err := databaseConnection(envConfig.DbUri)
	defer cancel()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Database Connection Error")
	}

	var (
		config = fiber.Config{
			ErrorHandler: api.ErrorHandler,
		}
		userService    = user.NewService(client, envConfig.DbName)
		sessionService = session.NewService(client, envConfig.DbName)

		userHandler    = api.NewUserHandler(userService)
		sessionHandler = api.NewSessionHandler(sessionService)

		app = fiber.New(config)
		//auth        = app.Group("/api")
		apiv1 = app.Group("/api/v1", api.Jwt())
	)

	api.UserRouter(apiv1, userHandler)
	api.SessionRouter(apiv1, sessionHandler)

	go func() {
		if err := app.Listen(*listenAddr); err != nil {
			log.Fatal().Err(err).Msg("cannot start http service")
		}
	}()

	ctxSig, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()
	<-ctxSig.Done()

	log.Info().Msg("got interruption signal")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(timeoutCtx); err != nil {
		log.Fatal().Err(err).Msg("failed stop http service")
	}

	log.Info().Msg("app is shutting down.")
}

func databaseConnection(dbConnStr string) (*mongo.Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnStr).SetServerSelectionTimeout(5*time.Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return client, cancel, nil
}
