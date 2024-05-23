package routes

import (
	"github.com/everest1508/carvia/handlers"
	"github.com/gin-gonic/gin"
)

func BlogRoute() {
	r := gin.Default()
	r.POST("/blog",handlers.PostBlog)
	r.GET("/blog",handlers.GetBlogs)
	r.GET("/blog/:id",handlers.GetBlog)
	r.Run(":8080")
}