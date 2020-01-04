package setting

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	confPath = "conf/.env"
)

func LoadConf() {
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Fatal error config file:%s \n", err)
	}
}
