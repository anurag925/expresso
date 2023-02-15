package configurations

import (
	env "github.com/caarlos0/env/v6"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

const (
	// EnvLocal represents the local Environment
	EnvLocal Environment = "local"

	// EnvTest represents the test Environment
	EnvTest Environment = "test"

	// EnvDevelop represents the development Environment
	EnvDevelopment Environment = "development"

	// EnvStaging represents the staging Environment
	EnvStaging Environment = "staging"

	// EnvQA represents the qa Environment
	EnvQA Environment = "qa"

	// EnvProduction represents the production Environment
	EnvProduction Environment = "production"
)

type app struct {
	ENV Environment `env:"ENV" envDefault:"development"`

	Dir                string `env:"DIR" envDefault:"${PWD}/../.."`
	TimeZone           string `env:"TIME_ZONE"`
	Timeout            int    `env:"Timeout"`
	Port               int    `env:"PORT"`
	KuveraJwtSecretKey string `env:"JWT_SECRET_KEY"`
	JwtSecretKey       string `env:"KURUKSHETRA_JWT_SECRET_KEY"`

	DBUserName string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	DBPort     int    `env:"DB_PORT"`
	DBHost     string `env:"DB_HOST"`

	MongoDBUrl  string `env:"MONGO_DB_URL"`
	MongoDBName string `env:"MONGO_DB_NAME"`

	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisPort     int    `env:"REDIS_PORT"`
	RedisHost     string `env:"REDIS_HOST"`
	RedisDBName   string `env:"REDIS_DB_NAME"`
	RedisPoolSize int    `env:"REDIS_POOL_SIZE"`
	RedisMaxRetry int    `env:"REDIS_MAX_RETRY"`
	RedisDB       int    `env:"REDIS_DB"`
}

type Configuration interface {
	Env() Environment
	Settings() *settings
	Secrets() *secrets
	App() *app
}

type Config struct {
	*settings
	*secrets
	*app
}

var _ Configuration = (*Config)(nil)

func (c *Config) Env() Environment {
	return c.App().ENV
}

func (c *Config) Settings() *settings {
	return c.settings
}

func (c *Config) Secrets() *secrets {
	return c.secrets
}

func (c *Config) App() *app {
	return c.app
}

func (dotenv *app) parseDotEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	if err := env.Parse(dotenv); err != nil {
		return err
	}
	return nil
}

func LoadConfigs() (*Config, error) {
	c := &Config{
		settings: &settings{},
		secrets:  &secrets{},
		app:      &app{},
	}
	if err := c.app.parseDotEnv(); err != nil {
		return c, err
	}
	// Load the config file
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings")
	viper.AddConfigPath(c.Dir + "/configs/configurations/")
	if err := viper.ReadInConfig(); err != nil {
		return c, err
	}

	if err := viper.Unmarshal(c.settings); err != nil {
		return c, err
	}
	viper.SetConfigName("secrets")
	viper.AddConfigPath(c.Dir + "/configs/configurations/")
	if err := viper.MergeInConfig(); err != nil {
		return c, err
	}
	if err := viper.Unmarshal(c.secrets); err != nil {
		return c, err
	}
	return c, nil
}

var DefaultConfigSet = wire.NewSet(
	LoadConfigs,
	wire.Bind(new(Configuration), new(*Config)),
)
