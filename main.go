package main

import (
	"github.com/everest1508/carvia/db"
	"github.com/everest1508/carvia/routes"
)

func init(){
	db.DBConnect()
}

func main(){
	routes.BlogRoute()
}