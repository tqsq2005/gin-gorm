package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Test()  {
	log.Println(viper.GetString("APP_NAME"))
}