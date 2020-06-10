package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const port = ":3000"

type response struct {
	Message int `json:"message"`
}

func main() {
	http.HandleFunc("/", reverseProxyHandler)

	log.Println("Starting proxy server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

// Idea from: https://www.integralist.co.uk/posts/golang-reverse-proxy/#3
func reverseProxyHandler(res http.ResponseWriter, req *http.Request) {
	origin, _ := url.Parse("http://localhost:3001/")

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(res, req)
}

func jsonResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
