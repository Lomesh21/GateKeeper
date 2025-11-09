package gateway

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyHandler(target string) (http.HandlerFunc, error) {
	backendURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request to %s", backendURL.String())
		proxy.ServeHTTP(w, r)
	}, nil
}
