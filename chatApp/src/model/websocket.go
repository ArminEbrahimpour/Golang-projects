package model

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"

	"github.com/gofiber/websocket/v2"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

func (ws *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// register a new client
	ws.clients[ctx] = true
	defer func() {
		delete(ws.clients, ctx)
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		// send the message to the broadcast chan
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println(err)
		}
		ws.broadcast <- &message
	}

}

func (ws *WebSocketServer) HandleMessages() {
	for {
		msg := <-ws.broadcast

		// send msg to all clients
		for client := range ws.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))

			if err != nil {
				log.Println(err)
				client.Close()
				delete(ws.clients, client)
			}
		}

	}
}

func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("views/messages.html")
	if err != nil {
		log.Fatalf("template parsing : %s", err)
	}

	// render the template
	var rendered_message bytes.Buffer
	err = tmpl.Execute(&rendered_message, msg)
	if err != nil {
		log.Fatalf("template exception : %s", err)
	}
	return rendered_message.Bytes()
}
