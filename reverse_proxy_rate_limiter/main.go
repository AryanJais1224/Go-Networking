package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type bucket struct {
	tokens     int
	lastRefill time.Time
}

var (
	rate       = 5
	capacity   = 10
	clients    = make(map[string]*bucket)
	mu         sync.Mutex
	target, _  = url.Parse("http://localhost:9000")
)

func allow(ip string) bool {
	mu.Lock()
	defer mu.Unlock()

	b, exists := clients[ip]
	if !exists {
		clients[ip] = &bucket{tokens: capacity, lastRefill: time.Now()}
		return true
	}

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	newTokens := int(elapsed * float64(rate))

	if newTokens > 0 {
		b.tokens += newTokens
		if b.tokens > capacity {
			b.tokens = capacity
		}
		b.lastRefill = now
	}

	if b.tokens > 0 {
		b.tokens--
		return true
	}

	return false
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr
	return ip
}

func main() {
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		if !allow(ip) {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
