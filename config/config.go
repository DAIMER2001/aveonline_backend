package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Host     string
		DBName   string
		Port     int
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$HOME", "Proyectos/go/golan-rest-simple", "config"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}