package config

type (
	Config struct {
		Main    Main    `envPrefix:"MAIN_" env:"inline"`
		Mongodb Mongodb `envPrefix:"MONGO_" env:"inline"`
		TMDB    TMDB    `envPrefix:"TMDB_" env:"inline"`
	}

	Main struct {
		ServiceName       string `env:"SERVICE_NAME" envDefault:"tmdb-parser"`
		LogLevel          string `env:"LOG_LEVEL" envDefault:"info"`
		LogForcePlainText bool   `env:"LOG_FORCE_PLAIN_TEXT" envDefault:"false"`
	}

	Mongodb struct {
		URL  string `env:"URL,required"`
		Name string `env:"DB_NAME,required"`
	}

	TMDB struct {
		APIKey string `env:"API_KEY,required"`
	}
)
