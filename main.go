/*
 * File: \main.go                                                              *
 * Project: blog_service                                                       *
 * Created At: Friday, 2022/05/27 , 20:47:27                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/06/1 , 20:33:08                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"go_start/blog_service/global"
	"go_start/blog_service/internals/model"
	"go_start/blog_service/internals/routers"
	"go_start/blog_service/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// @title blog service
// @version v0.0
// @description This is a simple blog service based on golang,mariadb and okteto
// @contact.name Elchn

// @host https://blog-service-elchn.cloud.okteto.net
// @BasePath /api/v1
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	routers := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        routers,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}
