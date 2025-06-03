package main

import "net/http"

type Message struct {
	ClientID string
	Text     string
}

type Hub struct {
	id        string
	clients   map[*Client]bool
	broadCast chan *Message
	register  chan *Client
}

func serveHub(hub Hub, w http.ResponseWriter, r *http.Request) {}
