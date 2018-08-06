package config

import (
	"github.com/spf13/viper"
)

// Config contains app cfg variables
type Config struct {
	Port             string
	PostgresDB       string
	PostgresHost     string
	PostgresPassword string
	PostgresPort     int
	PostgresUser     string
}

// New returns Config
func New() (*Config, error) {
	viper.SetDefault("PORT", "9090")
	viper.SetDefault("POSTGRESDB", "postgres")
	viper.SetDefault("POSTGRESHOST", "localhost")
	viper.SetDefault("POSTGRESPASSWORD", "postgres")
	viper.SetDefault("POSTGRESPORT", 5440)
	viper.SetDefault("POSTGRESUSER", "postgres")

	viper.AutomaticEnv()

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
