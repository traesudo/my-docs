package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-docs/api"
	"my-docs/api/docs"
	"my-docs/api/users"
)

func main() {
	r := gin.Default()
	r.Use(APIControllerMiddleware())
	controller := api.NewAPIController()
	// 缺少依赖注入
	indexApi := r.Group("api")
	docRunner := docs.DocsController{controller}
	indexApi.GET("posts", controller.Handler(docRunner.GetPosts))
	indexApi.GET("post/:id", controller.Handler(docRunner.GetPost))
	indexApi.GET("directions", controller.Handler(docRunner.GetDirections))
	indexApi.GET("search", controller.Handler(docRunner.GlobalSearch))
	indexApi.GET("lives", controller.Handler(docRunner.GlobalSearch))
	userApi := r.Group("user")
	//  缺少依赖注入
	useRunner := users.UserController{controller}
	userApi.POST("signup", controller.Handler(useRunner.SignUp))

	r.Run()
	log.Println("start run ...")
}

func APIControllerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller := &api.APIController{
			Context: c,
		}
		c.Set("controller", controller)
		c.Next()
	}
}
