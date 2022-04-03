package main

import (
	"context"
	config "food-delivery/configs"
	"food-delivery/internal/handler"
	"food-delivery/internal/repository"
	protorepository "food-delivery/internal/repository/proto"
	"food-delivery/internal/service"
	"food-delivery/internal/service/proto"
	"food-delivery/server"
	_ "github.com/golang-migrate/migrate"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Food-Delivery API
// @version 1.0
// @description REST API for Delivery

// @host 165.232.68.67:8081
// @BasePath /

func main() {
	cfg, err := config.Init("./configs")
	if err != nil {
		log.Fatal().Err(err).Msg("wrong config variables")
	}

	db, err := newPostgresDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("err initializing DB")
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo, cfg)
	handlers := handler.NewHandler(services)
	srv := server.NewServer(cfg, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); err != http.ErrServerClosed {
			log.Error().Err(err).Msg("error occurred while running http server")
		}
	}()

	repoGRPC := protorepository.NewRepository(db)
	servicesGRPC := proto.NewService(repoGRPC, cfg)
	srvGRPC := server.NewServerGRPC()
	srvGRPC.RegisterServices(servicesGRPC)

	go func() {
		if err := srvGRPC.Run(cfg); err != nil {
			log.Error().Err(err).Msg("error occurred while running gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to stop server")
	}

	if err := db.Close(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop connection db")
	}

}

func newPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	return repository.NewPostgresDB(repository.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Dbname,
		SSLMode:  cfg.Postgres.Sslmode,
	})
}
