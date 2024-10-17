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
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // Path to Config file
	viper.SetConfigName("local")      // Tên File Config
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w", err))
	}

	// fmt.Println("Server port: ", viper.GetInt("server.port"))
	// fmt.Println("Security port: ", viper.GetString("security.jwt.key"))

	//configuration
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Printf("Server port: %d \n", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database user: %s, password: %s, host: %s, dbName: %s \n", db.User, db.Password, db.Host, db.DbName)
	}
}
