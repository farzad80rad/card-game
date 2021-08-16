package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


func tempHanderl(c *gin.Context) {
	c.String(http.StatusOK, "Hello world" )
}

func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, content-type, method, x-pingother")
	c.Header("Access-Control-Allow-Credentials", "true")
	if strings.EqualFold("OPTIONS", c.Request.Method) {
		c.JSON(http.StatusOK, struct{}{})
		return
	}
	c.Next()
}



func Init() *gin.Engine{
	router := gin.Default()
	initSys()
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware)

	hokm := router.Group("/hokm")
	{
		hokm.POST("/playWithBot", handelPlayWithBot)
		hokm.POST("/playNormal",tempHanderl)
	}
	return router
}