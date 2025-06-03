package main

import (
	"net/http"

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

func (h *Hub) Run() {}

func serveHub(hub Hub, w http.ResponseWriter, r *http.Request) {}
