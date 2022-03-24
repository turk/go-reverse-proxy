package main

import (
        "net/http"
        "net/http/httputil"
        "net/url"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
        proxyURL, _ := url.Parse(target)
        proxy := httputil.NewSingleHostReverseProxy(proxyURL)
        req.URL.Scheme = proxyURL.Scheme
        req.Header.Set("X-Forwarded-Host", req.Host)

        proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
        var targetURL string

        switch req.Host {
        case "host1.localhost":
                targetURL = "http://localhost:1111"
        case "host2.localhost":
                targetURL = "http://localhost:2222"
        }

        serveReverseProxy(targetURL, res, req)
}

func main() {
        http.HandleFunc("/", handleRequestAndRedirect)
        if err := http.ListenAndServe(":80", nil); err != nil {
                panic(err)
        }
}
