package main

import (
	"backend/controllers/usercontroller"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/fetch-and-save-users", usercontroller.FetchAndSaveUsers)

	r.GET("/api/users", usercontroller.Index)
	r.GET("/api/user/:id", usercontroller.Show)
	r.POST("/api/user", usercontroller.Create)
	r.PUT("/api/users", usercontroller.Update)
	r.DELETE("/api/users", usercontroller.Delete)

	r.Run()
}
