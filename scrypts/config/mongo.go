package config

import (
	"github.com/caarlos0/env/v8"
	"go.uber.org/zap"
)

type Mongo struct {
	DB  string `env:"MONGO_DB" envDefault:"kp"`
	URI string `env:"MONGO_URI,notEmpty"`
}

func NewMongoConfig() *Mongo {
	logger := zap.L()
	conf := &Mongo{}

	if err := env.Parse(conf); err != nil {
		logger.Panic("Failed to parse Mongo config", zap.Error(err))
		return nil
	}

	return conf
}
