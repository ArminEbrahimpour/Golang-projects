package main

import (
	"bytes"
	"log"
	"text/template"

	"github.com/google/uuid"
)

type Message struct {
	ClientID uuid.UUID
	Text     string
}

type WsMessage struct {
	Text    string      `json:"text"`
	Headers interface{} `json:"headers"`
}

type Hub struct {
	clients    map[*Client]bool
	messages   []*Message
	broadCast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadCast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {

	for {

		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("client registered %s", client.id)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}
		case msg := <-h.broadCast:
			h.messages = append(h.messages, msg)

			for client := range h.clients {
				select {
				case client.send <- getMessageTemplate(msg):
				default:
					close(client.send)
					delete(h.clients, client)

				}
			}
		}

	}

}
func getMessageTemplate(msg *Message) []byte {

	tmpl, err := template.ParseFiles("templates/message.html")
	if err != nil {
		log.Println(err)
	}
	var renderedMessage bytes.Buffer

	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Println(err)
	}
	return renderedMessage.Bytes()
}
