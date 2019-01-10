package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"Kiddy/routers"
	"Kiddy/setting"
	"net/http"
)

func main() {
	gin.SetMode(setting.RunMode)
	router:=routers.InitRouters()
	s := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
