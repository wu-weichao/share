package routers

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
	"share/internal/app/middleware"
	v1 "share/internal/app/service/api/v1"
	"share/internal/app/service/web"
	"share/pkg/html/funcs"
	"strings"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()

	// middleware
	r.Use(middleware.Cors())

	// templates
	r.HTMLRender = loadTemplates("../../web/template/blog")

	// static
	r.StaticFS("/static", http.Dir("../../web/static"))

	// router
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/blog")
	})

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
			apiv1WithAuth.DELETE("/tags/:id", v1.DeleteTag)

			apiv1WithAuth.GET("/articles", v1.GetArticles)
			apiv1WithAuth.GET("/articles/:id", v1.GetArticle)
			apiv1WithAuth.POST("/articles", v1.StoreArticle)
			apiv1WithAuth.PUT("/articles/:id", v1.UpdateArticle)
			apiv1WithAuth.DELETE("/articles/:id", v1.DeleteArticle)
			apiv1WithAuth.PUT("/articles/:id/publish", v1.PublishArticle)
			apiv1WithAuth.PUT("/articles/:id/unpublish", v1.UnpublishArticle)

			apiv1WithAuth.GET("/topics", v1.GetTopics)
			apiv1WithAuth.GET("/topics/:id", v1.GetTopic)
			apiv1WithAuth.POST("/topics", v1.StoreTopic)
			apiv1WithAuth.PUT("/topics/:id", v1.UpdateTopic)
			apiv1WithAuth.DELETE("/topics/:id", v1.DeleteTopic)
		}
	}
	// frontend group
	front := r.Group("/blog")
	{
		front.GET("/", web.Homepage)
		front.GET("/page/:num", web.Homepage)
		front.GET("/article/:id", web.ShowArticle)
		front.GET("/topic/:flag", web.Topic)
		front.GET("/topic/:flag/page/:num", web.Topic)
	}

	return r
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// template functions
	funcMap := template.FuncMap{
		"unixToDateTime": funcs.UnixToDateTime,
		"unixToDate":     funcs.UnixToDate,
		"unixToFormat":   funcs.UnixToFormat,
		"implode":        funcs.Implode,
		"html":           funcs.Html,
	}

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(fmt.Sprintf("load template.layouts err: %v\n", err))
	}

	components, err := filepath.Glob(templatesDir + "/components/*.tmpl")
	if err != nil {
		panic(fmt.Sprintf("load template.components err: %v\n", err))
	}

	views, err := filepath.Glob(templatesDir + "/views/*.tmpl")
	if err != nil {
		panic(fmt.Sprintf("load template.views err: %v\n", err))
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, view := range views {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		// append components
		for _, component := range components {
			layoutCopy = append(layoutCopy, component)
		}
		files := append(layoutCopy, view)
		r.AddFromFilesFuncs(strings.TrimSuffix(filepath.Base(view), filepath.Ext(view)), funcMap, files...)
	}
	return r
}
