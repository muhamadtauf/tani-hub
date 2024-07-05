package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tani-hub/controller"
	"tani-hub/middleware"
)

func NewRouter(tagController *controller.TagController, categoryController *controller.CategoryController, articleController *controller.ArticleController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//api group
	router := service.Group("/api")

	//auth route
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/validate", middleware.RequireAuth, controller.Validate)

	//tags route
	tagRouter := router.Group("/tag")
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	//category route
	categoryRouter := router.Group("/category")
	categoryRouter.GET("", categoryController.FindAll)
	categoryRouter.GET("/:categoryId", categoryController.FindById)
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PUT("/:categoryId", categoryController.Update)
	categoryRouter.DELETE("/:categoryId", categoryController.Delete)

	//article route
	articleRouter := router.Group("/article")
	articleRouter.GET("", articleController.FindAll)
	articleRouter.GET("/:articleId", articleController.FindById)
	articleRouter.POST("", articleController.Create)
	articleRouter.PUT("/:articleId", articleController.Update)
	articleRouter.DELETE("/:articleId", articleController.Delete)

	return service
}
