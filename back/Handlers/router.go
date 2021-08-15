package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func tempHanderl(c *gin.Context) {
	c.String(http.StatusOK, "Hello world" )
}

func Start(){
	router := gin.Default()
	router.GET("/someGet", tempHanderl)
	router.POST("/somePost", tempHanderl)
	router.PUT("/somePut", tempHanderl)
	router.DELETE("/someDelete", tempHanderl)
	router.PATCH("/somePatch", tempHanderl)
	router.HEAD("/someHead", tempHanderl)
	router.OPTIONS("/someOptions", tempHanderl)
}