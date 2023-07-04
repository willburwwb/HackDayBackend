package main

import (
	"HackDayBackend/configs"
	"HackDayBackend/global"
	"HackDayBackend/internal/router"
	"HackDayBackend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := global.Set(); err != nil {
		panic(err)
	}

	r := router.NewRouter()
	gin.SetMode(configs.Server_config.RunMode)
	s := &http.Server{
		Addr:         configs.Server_config.Addr,
		Handler:      r,
		ReadTimeout:  configs.Server_config.ReadTimeout * time.Second,
		WriteTimeout: configs.Server_config.WriteTimeout * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		utils.DebugF("listen err: %s", err)
		panic(err)
	}
}
