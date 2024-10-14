package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
	Security struct {
		Jwt struct {
			Secret string `mapstructure:"secret"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	serverPort := viper.GetInt("server.port")
	jwtSecret := viper.GetString("security.jwt.secret")
	fmt.Println("Server port: ", serverPort)
	fmt.Println("JWT secret: ", jwtSecret)

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}

	fmt.Println("Server port: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database: user=%s, password=%s, host=%s, dbName=%s\n", db.User, db.Password, db.Host, db.DbName)
	}
}
