package main

//import "github.com/gin-gonic/gin"
import (
	"fmt"
	"share/configs"
	"share/internal/app/routers"
)

func main() {
	r := routers.InitRouter()
	fmt.Printf("configs.App: %v+\n", configs.App)
	r.Run()
}
