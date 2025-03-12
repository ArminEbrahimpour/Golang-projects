package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ArminEbrahimpour/caching-proxy/internal/proxy"
)

func main() {

	var p *proxy.Proxy

	port := flag.Int("port", 1234, "declare the port")

	Origin := flag.String("origin", "something", "put the origin site in here")

	flag.Parse()
	p = proxy.NewProxy(*Origin)

	http.HandleFunc("/", p.ServeHttp)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), nil))
}
