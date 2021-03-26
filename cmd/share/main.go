package main

//import "github.com/gin-gonic/gin"
import (
	"fmt"
	"share/configs"
	"share/internal/app/routers"
	_ "share/internal/database"
	_ "share/internal/models"
)

func main() {
	r := routers.InitRouter()
	fmt.Printf("configs.App: %+v\n", configs.App)
	fmt.Printf("configs.Database: %+v\n", configs.Database)
	r.Run()
}
