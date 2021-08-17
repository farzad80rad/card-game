package Handlers

import (
	"fmt"
	"log"
	"net/http"
)

func InitWebSoucketRouter(){
	http.HandleFunc("/hokm/websocketBot", addNewWebSocket_B )
	http.HandleFunc("/tss", func(w http.ResponseWriter, r * http.Request){ fmt.Fprintln(w, "tsssss") } )
	log.Println(http.ListenAndServe(":8081",nil))
}