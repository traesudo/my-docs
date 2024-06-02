package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-docs/api"
)

func main() {
	r := gin.Default()
	indexapi := r.Group("api")
	indexapi.GET("posts", api.GetPosts)
	indexapi.GET("post/:id", api.GetPost)
	indexapi.GET("directions", api.GetDirections)
	r.Run()
	log.Println("start run ...")
}
