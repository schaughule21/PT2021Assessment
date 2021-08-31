package main

import (
	"31Aug-Assessment/database"
	"31Aug-Assessment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/addfactory", func(c *gin.Context) {
		var factory models.Factory
		err := c.Bind(&factory)
		database.CheckError(err)
		err = database.CreateFactory(factory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "Factory Details Added")
		}
	})
	r.GET("/getfactory", func(c *gin.Context) {
		res := database.GetFactory()
		c.JSON(200, res)
	})
	r.POST("/login", Login)
	return r
}
