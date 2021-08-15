package main

import (
	"github.com/farzad80rad/cards/back/Handlers"
)

func main() {
	router := Handlers.Init()
	router.Run();
}
