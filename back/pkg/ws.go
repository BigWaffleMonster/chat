package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Event    string    `json:"event"`
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Body     string    `json:"body"`
}

type Server struct {
	clients       map[*websocket.Conn]bool
	handleMessage func(msg Message)
}

func StartServer(handleMessage func(msg Message)) *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		handleMessage,
	}

	http.HandleFunc("/", server.chat)

	go http.ListenAndServe(":8080", nil)

	forever := make(chan bool)
	log.Printf("[*] Server has been started on port 8080. To exit press CTRL+C")
	<-forever

	return &server
}

func (s *Server) chat(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	s.clients[connection] = true
	defer delete(s.clients, connection)

	for {
		var msg Message
		err := connection.ReadJSON(&msg)
		if err != nil || msg.Event == "close" {
			break
		}

		switch msg.Event {
		case "connection":
			s.Broadcast(msg)
		case "message":
			s.Broadcast(msg)
		}

		go s.handleMessage(msg)
	}
}

func (s *Server) Broadcast(msg Message) {
	for conn := range s.clients {
		conn.WriteJSON(msg)
	}
}
