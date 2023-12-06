package main

import (
	"fmt"
	"net/http"

	"github.com/CelepiYakup/golang-chat/pkg/websocket"
)

func setupRoutes() {
	pool := websocket.NewPool()
}

func main() {
	fmt.Println("Yakup's full stack chat project")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
