package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func serveWebsock(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// upgrade the connection
}
