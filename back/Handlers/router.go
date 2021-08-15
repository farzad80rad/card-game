package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func tempHanderl(c *gin.Context) {
	c.String(http.StatusOK, "Hello world" )
}


func Init() *gin.Engine{
	router := gin.Default()
	initSys()

	hokm := router.Group("/hokm")
	{
		hokm.POST("/hokm/enter", tempHanderl)
	}
	return router
}