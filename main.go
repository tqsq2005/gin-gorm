package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tqsq2005/gin-gorm/routers"
	"github.com/tqsq2005/gin-gorm/setting"
	"net/http"
	"time"
)

func init() {
	setting.LoadConf()
}

func main() {
	router			:= routers.InitRouter()
	readTimeout 	:= viper.GetDuration("APP_READ_TIMEOUT") * time.Second
	writeTimeout 	:= viper.GetDuration("APP_WRITE_TIMEOUT") * time.Second
	endPoint 		:= fmt.Sprintf(":%d", viper.GetInt("APP_HTTP_PORT"))
	maxHeaderBytes 	:= 1 << 20

	s := &http.Server{
		Addr:			endPoint,
		Handler:		router,
		ReadTimeout:	readTimeout,
		WriteTimeout:	writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("Start http server listening %s", endPoint)

	log.Println(s.ListenAndServe())
}
