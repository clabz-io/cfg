package cfg

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ApiSecret       string `mapstructure:"API_SECRET"`
	DBHost          string `mapstructure:"MYSQL_HOST"`
	DBName          string `mapstructure:"MYSQL_DB"`
	DBNameTest      string `mapstructure:"MYSQL_DB_TEST"`
	DBPassword      string `mapstructure:"MYSQL_PASS"`
	DBPort          string `mapstructure:"MYSQL_PORT"`
	DBUrl           string `mapstructure:"MYSQL_DB_URL"`
	AppDomain       string `mapstructure:"APP_DOMAIN"`
	DBUsername      string `mapstructure:"MYSQL_USER"`
	Environment     string `mapstructure:"ENVIRONMENT"`
	Host            string `mapstructure:"HOST"`
	Port            string `mapstructure:"PORT"`
	TypesenseAPIKEY string `mapstructure:"TYPESENSE_APIKEY"`
	TypesenseAddr   string `mapstructure:"TYPESENSE_ADDR"`

	AccessTokenDuration  time.Duration `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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
