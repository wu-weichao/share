package routers

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	//"share/internal/app/middleware"
	"share/internal/app/service/api/v1"
	"share/internal/app/service/web"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()

	// middleware
	r.Use()

	// templates
	r.HTMLRender = loadTemplates("../../web/template/blog")

	// static
	r.StaticFS("/static", http.Dir("../../web/static"))

	// router
	r.GET("/", web.Homepage)

	// route group
	// api group
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/login", v1.Login)
		//apiv1.POST("/logout")

		// with auth token
		//apiv1WithAuth := apiv1.Group("")
		//apiv1WithAuth.Use(middleware.JwtCheck())
		{
			//apiv1WithAuth.GET("/user_info")
		}
	}
	// frontend group
	front := r.Group("/blog")
	{
		front.GET("/", web.Homepage)
	}

	return r
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(fmt.Sprintf("load template.layouts err: %v\n", err))
	}

	views, err := filepath.Glob(templatesDir + "/views/*.tmpl")
	if err != nil {
		panic(fmt.Sprintf("load template.views err: %v\n", err))
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, view := range views {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, view)
		r.AddFromFiles(filepath.Base(view), files...)
	}
	return r
}
