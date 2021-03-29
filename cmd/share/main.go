package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"share/configs"
	"share/internal/app/routers"
	_ "share/internal/database"
	_ "share/internal/models"
	_ "share/internal/schedule"
)

func main() {
	// set gin mode
	gin.SetMode(configs.App.RunMode)
	// init router
	r := routers.InitRouter()
	// run http server
	addr := fmt.Sprintf(":%d", configs.App.HttpPort)
	r.Run(addr)
	log.Printf("[info] start http server listening %s", addr)
}
