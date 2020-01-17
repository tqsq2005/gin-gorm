package main

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tqsq2005/gin-gorm/pkg/logging"
	"github.com/tqsq2005/gin-gorm/routers"
	"github.com/tqsq2005/gin-gorm/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	setting.LoadConf()
	filePath := logging.GetLogFileFullPath()
	F := logging.OpenLogFile(filePath)
	log.SetOutput(F)
}

// @title API列表
// @version 1.0
// @description Gin-Gorm的API列表.
// @termsOfService http://github.com/tqsq2005/gin-gorm

// @contact.name API Support
// @contact.url http://github.com/tqsq2005/gin-gorm
// @contact.email tqsq2005@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	endLess()
}

func endLess()  {
	router						 := routers.InitRouter()
	endless.DefaultReadTimeOut    = viper.GetDuration("APP_READ_TIMEOUT") * time.Second
	endless.DefaultWriteTimeOut   = viper.GetDuration("APP_WRITE_TIMEOUT") * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint 					 := fmt.Sprintf(":%d", viper.GetInt("APP_HTTP_PORT"))

	server := endless.NewServer(endPoint, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	log.Printf("Start http server listening %s", endPoint)

	log.Println(server.ListenAndServe())
}

func shutDown()  {
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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
