package main

import (
	"github.com/everest1508/carvia/db"
	"github.com/everest1508/carvia/models"
)

func init() {
	db.DBConnect()
}

func main() {
	db.DB.AutoMigrate(&models.Blog{})
}