package routers

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"share/internal/app/middleware"
	v1 "share/internal/app/service/api/v1"

	"share/internal/app/service/web"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()

	// middleware
	r.Use(middleware.Cors())

	// templates
	r.HTMLRender = loadTemplates("../../web/template/blog")
	//r.SetFuncMap()

	// static
	r.StaticFS("/static", http.Dir("../../web/static"))

	// router
	r.GET("/", web.Homepage)

	// route group
	// api group
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/logout", v1.Logout)

		// with auth token
		apiv1WithAuth := apiv1.Group("")
		apiv1WithAuth.Use(middleware.JWTAuth())
		{
			apiv1WithAuth.GET("/user_info", v1.LoginUserInfo)

			apiv1WithAuth.GET("/tags", v1.GetTags)
			apiv1WithAuth.GET("/simple_tags", v1.GetSimpleTags)
			apiv1WithAuth.GET("/tags/:id", v1.GetTag)
			apiv1WithAuth.POST("/tags", v1.StoreTag)
			apiv1WithAuth.PUT("/tags/:id", v1.UpdateTag)
			apiv1WithAuth.DELETE("/tags/:id")

			apiv1WithAuth.GET("/articles", v1.GetArticles)
			apiv1WithAuth.GET("/articles/:id", v1.GetArticle)
			apiv1WithAuth.POST("/articles", v1.StoreArticle)
			apiv1WithAuth.PUT("/articles/:id", v1.UpdateArticle)
			apiv1WithAuth.DELETE("/articles/:id")
			apiv1WithAuth.PUT("/articles/:id/publish", v1.PublishArticle)
			apiv1WithAuth.PUT("/articles/:id/unpublish", v1.UnpublishArticle)
		}
	}
	// frontend group
	front := r.Group("/blog")
	{
		front.GET("/", web.Homepage)
		front.GET("/page/:num", web.Homepage)
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
