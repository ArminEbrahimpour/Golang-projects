package cache

import (
	"net/http"
	"time"
)

type Cache_object struct {
	Response *http.Response
	Content  []byte
	Time     time.Time
	TTL      time.Duration
}
