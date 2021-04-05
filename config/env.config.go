package config

import "github.com/spf13/viper"

type Config struct {
	DbHost string `mapstructure:"DB_HOST"`
	DbUser string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbDatabaseName string `mapstructure:"DB_DATABASE_NAME"`
	DbPort string `mapstructure:"DB_PORT"`

	AuthAccessTokenSecret  string `mapstructure:"AUTH_ACCESS_TOKEN_SECRET"`
	AuthAccessTokenExp     uint   `mapstructure:"AUTH_ACCESS_TOKEN_EXP"`
	AuthRefreshTokenSecret string `mapstructure:"AUTH_REFRESH_TOKEN_SECRET"`
	AuthRefreshTokenExp    uint   `mapstructure:"AUTH_REFRESH_TOKEN_EXP"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
    if err != nil {
        return
    }

	err = viper.Unmarshal(&config)
	return
}