package main

import (
	"log"
	"net/http"

	"github.com/ArminEbrahimpour/caching-proxy/internal/proxy"
)

func main() {

	var p *proxy.Proxy
	Origin := "https://basalam.ir"
	p = proxy.NewProxy(Origin)

	http.HandleFunc("/", p.ServeHttp)

	log.Fatal(http.ListenAndServe("127.0.0.1:6266", nil))
}
