package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresDB *Database
	HTTP       *Endpoint

	UserService    *Endpoint
	ProductService *Endpoint
	CouponService  *Endpoint
	GatewayService *Endpoint

	SymetricKey string
}

type config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBDatabase string `mapstructure:"DB_NAME"`

	UserGRPCHost string `mapstructure:"USER_GRPC_HOST"`
	UserGRPCPort string `mapstructure:"USER_GRPC_PORT"`

	ProductGRPCHost string `mapstructure:"PRODUCT_GRPC_HOST"`
	ProductGRPCPort string `mapstructure:"PRODUCT_GRPC_PORT"`

	CouponGRPCHost string `mapstructure:"COUPON_GRPC_HOST"`
	CouponGRPCPort string `mapstructure:"COUPON_GRPC_PORT"`

	GatewayGRPCHost string `mapstructure:"GATEWAY_GRPC_HOST"`
	GatewayGRPCPort string `mapstructure:"GATEWAY_GRPC_PORT"`

	SymetricKey string `mapstructure:"SYMETRIC_KEY"`
}

func LoadConfig(path string, service, env string) (*Config, error) {
	var cfg config
	viper.AddConfigPath(path)
	viper.SetConfigName(fmt.Sprintf("%s.%s", service, env))
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config file: %w", err)
	}

	return &Config{
		PostgresDB: &Database{
			Host:     cfg.DBHost,
			Port:     cfg.DBPort,
			User:     cfg.DBUser,
			Password: cfg.DBPassword,
			Database: cfg.DBDatabase,
		},
		UserService: &Endpoint{
			Host: cfg.UserGRPCHost,
			Port: cfg.UserGRPCPort,
		},

		ProductService: &Endpoint{
			Host: cfg.ProductGRPCHost,
			Port: cfg.ProductGRPCPort,
		},

		CouponService: &Endpoint{
			Host: cfg.CouponGRPCHost,
			Port: cfg.CouponGRPCPort,
		},
		GatewayService: &Endpoint{
			Host: cfg.GatewayGRPCHost,
			Port: cfg.GatewayGRPCPort,
		},
		SymetricKey: cfg.SymetricKey,
	}, nil
}
