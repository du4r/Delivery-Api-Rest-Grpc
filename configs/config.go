package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	APIHTTP APIConfig
	APIGRPC APIGrpc
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type APIGrpc struct{
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Init() {
	viper.SetDefault("api.port", 9000)
	viper.SetDefault("grpc.port", 50051)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil{
		if _ , ok := err.(viper.ConfigFileNotFoundError); !ok{
			return err
		}
	}
	
	cfg = new(config)

	cfg.APIHTTP = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.APIGRPC = APIGrpc{
		Port: viper.GetString("grpc.port"),
	}

	cfg.DB = DBConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDb() DBConfig{
	return cfg.DB
}

func GetHttpServerPort() string {
	return cfg.APIHTTP.Port
}

func GetGrpcServerPort() string {
	return cfg.APIGRPC.Port
}