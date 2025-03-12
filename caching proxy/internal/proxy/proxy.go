package proxy

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ArminEbrahimpour/caching-proxy/internal/cache"
)

type Proxy struct {
	Cache  map[string]*cache.Cache_object
	Origin string
	Mutex  sync.RWMutex
}

func NewProxy(origin string) *Proxy {
	return &Proxy{
		Origin: origin,
		Cache:  make(map[string]*cache.Cache_object),
	}
}

func (p *Proxy) ClearCash() {
	p.Mutex.Lock()
	p.Cache = make(map[string]*cache.Cache_object)
	fmt.Println("cache cleared ")
	p.Mutex.Unlock()
}

func (p *Proxy) ServeHttp(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	cache_key := r.Method + "$$" + r.URL.String()
	hash := md5.New()
	_, err := io.WriteString(hash, cache_key)
	if err != nil {
		panic(err)
	}
	hashed_key := hex.EncodeToString(hash.Sum(nil))
	p.Mutex.RLock()

	// if file is saved in cache
	if c, ok := p.Cache[hashed_key]; ok {
		p.Mutex.RUnlock()
		SetHeadersAndRespond(w, *c.Response, c.Content, "HIT", hashed_key)
		return
	}
	p.Mutex.RUnlock()

	// if file is not saved

	Origin_uri := p.Origin + r.URL.String()
	resp, err := http.Get(Origin_uri)

	if err != nil {
		http.Error(w, "Error forwarding the Request", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error forwarding request body ", http.StatusInternalServerError)
		return
	}

	// caching the response
	p.Mutex.Lock()
	p.Cache[hashed_key] = &cache.Cache_object{

		Response: resp,
		Content:  body,
		Time:     time.Now(),
	}
	p.Mutex.Unlock()
	SetHeadersAndRespond(w, *resp, body, "MISS", hashed_key)
}
func SetHeadersAndRespond(w http.ResponseWriter, resp http.Response, content []byte, HeaderValue string, KEY string) {
	w.Header().Set("X-Cache", HeaderValue)
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.Write(content)

}

func (p *Proxy) CleanExpiredCache() {

	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	for i, c := range p.Cache {
		if time.Since(time.Now()) > c.TTL {
			delete(p.Cache, i)
		}
	}

}
