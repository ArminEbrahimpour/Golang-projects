package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id   uuid.UUID
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

const (
	pongWait := 60 * time.Second

)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWebsock(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	id := uuid.New()
	client := &Client{id: id, hub: hub, conn: conn, send: make(chan []byte)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) writePump() {

	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

}

func (c *Client) readPump() {}
