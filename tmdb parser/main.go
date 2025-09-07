package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
	"tmdb_parser/config"
	"tmdb_parser/db/repository"
	"tmdb_parser/parser"
	"tmdb_parser/pkg/logger"
	"tmdb_parser/pkg/tmdb"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		log.Fatal().Err(err).Msg("Can't parse config")
	}

	rootLogger := logger.InitRootLogger(
		cfg.Main.LogForcePlainText,
		logger.ParseEnvLoggerEnv(cfg.Main.LogLevel),
		cfg.Main.ServiceName,
	)

	rootLogger.Info().Msg("Successfully connected to database")

	client, err := setupDatabaseConnection(ctx, cfg.Mongodb.URL)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't setup database connection")
	}
	defer client.Disconnect(ctx)
	mongoConnect := client.Database(cfg.Mongodb.Name)

	rootLogger.Info().Msg("Successfully connected")

	tmdbClient := tmdb.NewClient(cfg.TMDB.APIKey)

	parser := parser.New(repository.NewMovieRepository(mongoConnect), tmdbClient)
	parser.Run(ctx)
}
