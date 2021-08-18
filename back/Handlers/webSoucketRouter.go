package Handlers

import (
	"log"
	"net/http"
)

func InitWebSoucketRouter(){
	http.HandleFunc("/hokm/websocketBot", addNewWebSocket_B )
	log.Println(http.ListenAndServe(":8081",nil))
}