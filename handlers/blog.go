package handlers

import (
	"github.com/everest1508/carvia/db"
	"github.com/everest1508/carvia/models"
	"github.com/gin-gonic/gin"
)

func PostBlog(context *gin.Context) {
	var blog models.Blog
	context.BindJSON(&blog)
	db.DB.Create(&blog)
	context.JSON(201,gin.H{"title":blog.Title})
}

func GetBlogs(context *gin.Context) {
	var blog []models.Blog
	db.DB.Find(&blog)
	context.JSON(200,blog)
}

func GetBlog(context *gin.Context) {
	var blog models.Blog
	blog_id := context.Param("id")
	db.DB.First(&blog,blog_id)
	context.JSON(200,blog)
}