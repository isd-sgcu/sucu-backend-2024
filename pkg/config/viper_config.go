package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

type viperConfig struct {
	Server `mapstructure:",squash"`
	Db     `mapstructure:",squash"`
	Jwt    `mapstructure:",squash"`
	Aws    `mapstructure:",squash"`
}

var (
	once     sync.Once
	instance Config
)

func NewViperConfig() Config {
	once.Do(func() {
		serverEnv := os.Getenv("SERVER_ENV")
		if serverEnv == "production" {
			loadProductionEnv()
		} else {
			loadDefaultEnv()
		}
	})
	return instance
}

func loadProductionEnv() {
	log.Println("Loading production environment variables")
	instance = &viperConfig{
		Server: Server{
			Name: os.Getenv("SERVER_NAME"),
			Env:  os.Getenv("SERVER_ENV"),
			Url:  os.Getenv("SERVER_URL"),
			Host: os.Getenv("SERVER_HOST"),
			Port: func() int {
				port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
				if err != nil {
					panic("error while loading server port")
				}
				return port
			}(),
		},
		Db: Db{
			Host: os.Getenv("DB_HOST"),
			Port: func() int {
				port, err := strconv.Atoi(os.Getenv("DB_PORT"))
				if err != nil {
					panic("error while loading db port")
				}
				return port
			}(),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Timezone: os.Getenv("DB_TIMEZONE"),
		},
		Jwt: Jwt{
			ApiSecretKey:       os.Getenv("JWT_API_SECRET_KEY"),
			AccessTokenSecret:  os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			RefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
			AccessTokenExpiration: func() int {
				expiration, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION"))
				if err != nil {
					panic("error while loading access token expiration")
				}
				return expiration
			}(),
			RefreshTokenExpiration: func() int {
				expiration, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION"))
				if err != nil {
					panic("error while loading refresh token expiration")
				}
				return expiration
			}(),
		},
		Aws: Aws{
			BucketName:      os.Getenv("AWS_BUCKET_NAME"),
			AccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:          os.Getenv("AWS_REGION"),
		},
	}
}

func loadDefaultEnv() {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configs file: %s", err)
	}

	cfg := &viperConfig{}
	if err := v.Unmarshal(cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	instance = cfg
}

func GetConfig() Config {
	if instance == nil {
		instance = NewViperConfig()
	}
	return instance
}

func (c *viperConfig) GetServer() Server {
	return c.Server
}

func (c *viperConfig) GetDb() Db {
	return c.Db
}

func (c *viperConfig) GetJwt() Jwt {
	return c.Jwt
}

func (c *viperConfig) GetAws() Aws {
	return c.Aws
}
