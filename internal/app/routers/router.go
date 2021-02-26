package routers

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"share/internal/app/service/web"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// templates
	r.HTMLRender = loadTemplates("../../web/template/blog")

	// static
	r.StaticFS("/static", http.Dir("../../web/static"))

	// router
	r.GET("/", web.Homepage)

	// route group
	// api group
	v1 := r.Group("/api/v1")
	{
		v1.GET("/login")
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
