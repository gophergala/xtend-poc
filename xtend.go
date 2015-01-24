package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var users []Player

type Player struct {
	Name string `json:"name"`
	Conn *websocket.Conn
}

func Login(ws *websocket.Conn) {
	var player Player
	fmt.Println(websocket.JSON.Receive(ws, &player))
	player.Conn = ws
	users = append(users, player)

	if (len(users)) == 2 {
		websocket.JSON.Send(users[0].Conn, users[1].Name)
		websocket.JSON.Send(users[1].Conn, users[0].Name)
	}
	fmt.Println(player)
}

func main() {
	http.Handle("/login", websocket.Handler(Login))
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_PATH"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Listen on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
