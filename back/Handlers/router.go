package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func tempHanderl(c *gin.Context) {
	c.String(http.StatusOK, "Hello world" )
}

func InitRouter(){
	router := gin.Default()
	hokm := router.Group("/hokm")
	{
		hokm.POST("/hokm/enter", tempHanderl)
	}

}