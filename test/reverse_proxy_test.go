package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	backendService := "https://jandan.net"
	rpURL, err := url.Parse(backendService)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(rpURL)
	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		director(r)
		//r.Host = "target.local"
	}
	http.ListenAndServe(":8080", proxy)
}

var (
	// 建立域名和目标map
	hostTarget = map[string]string{
		"/false": "http://false.run/",
		"/p31": "https://www.p31.net/",
	}
	// 用于缓存 httputil.ReverseProxy
	hostProxy map[string]*httputil.ReverseProxy
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.RequestURI

	// 直接从缓存取出
	if fn, ok := hostProxy[host]; ok {
		r.RequestURI = strings.ReplaceAll(r.RequestURI, host, "")
		fn.ServeHTTP(w, r)
		return
	}

	// 检查域名白名单
	if target, ok := hostTarget[host]; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		if hostProxy == nil {
			hostProxy = make(map[string]*httputil.ReverseProxy)
		}
		hostProxy[host] = proxy // 放入缓存
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func TestProxy2(t *testing.T) {
	h := &baseHandle{}
	http.Handle("/", h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}

func TestProxy3(t *testing.T) {
	url1, _ := url.Parse("http://false.run")
	proxy := httputil.NewSingleHostReverseProxy(url1)
	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		fmt.Println(r)
		director(r)
	}
}
