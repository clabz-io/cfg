package cfg

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`

	DBUsername string `mapstructure:"MYSQL_USER"`
	DBPassword string `mapstructure:"MYSQL_PASS"`
	DBHost     string `mapstructure:"MYSQL_HOST"`
	DBPort     string `mapstructure:"MYSQL_PORT"`
	DBName     string `mapstructure:"MYSQL_DB"`
	DBNameTest string `mapstructure:"MYSQL_DB_TEST"`
	DBUrl      string

	ApiSecret           string        `mapstructure:"API_SECRET"`
	AccessTokenDuration time.Duration `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	// RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	TypesenseAddr   string `mapstructure:"TYPESENSE_ADDR"`
	TypesenseAPIKEY string `mapstructure:"TYPESENSE_APIKEY"`
}

func LoadConfig(name string, path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config: %v", err)
		return
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("config = %+v\n", config)
		log.Fatalf("config: %v", err)
		return
	}
	return
}
