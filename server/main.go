package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"unicode/utf8"
)

type Request struct {
	Message string `json:"message"`
}

const MaxMessageLength = 140

var wsUpgrader = websocket.Upgrader{}

var clients = make(map[*websocket.Conn]bool)

var broadcaster = make(chan Request)

func main() {
	go startBroadcasting()

	http.HandleFunc("/ws", ws)

	fmt.Println("WebSocket is working on http://localhost:4000/ws")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %+v\n", err)
		return
	}

	defer c.Close()
	defer delete(clients, c)

	clients[c] = true

	hello := Request{
		Message: fmt.Sprintf("[運営ちゃん] 新しい参加者が現れましたっ。これで%d人ですっ！", len(clients)),
	}
	broadcaster <- hello

	for {
		request := Request{}
		err := c.ReadJSON(&request)

		if err != nil {
			log.Printf("Error: %+v\n", err)

			bye := Request{
				Message: "[運営ちゃん] 誰かが退出したようです。。。",
			}
			broadcaster <- bye

			break
		}

		broadcaster <- request
	}
}

func startBroadcasting() {
	for {
		request := <-broadcaster

		if utf8.RuneCountInString(request.Message) > MaxMessageLength {
			continue
		}

		for client := range clients {
			err := client.WriteJSON(request)

			if err != nil {
				log.Printf("Error: %+v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
