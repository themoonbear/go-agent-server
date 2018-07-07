package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"flag"
	"fmt"
	"os"
	"strings"
	"log"
)

type handle struct {
	address string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	remote, err := url.Parse(h.address)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main()  {
	var address, port string
	flag.StringVar(&port, "p", "", "port")
	flag.StringVar(&address, "r", "", "remote address")
	flag.Parse()
	if address == "" || port == ""{
		fmt.Fprintf(os.Stderr, "Fatal error:%s", "参数为空\n")
		os.Exit(1)
	}
	if strings.HasSuffix(address, "/"){
		address = address[:len(address)-1]
	}
	h := &handle{address}
	fmt.Printf("agent server listen: %s\n", port)
	err := http.ListenAndServe(":" + port, h)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
		os.Exit(1)
	}
}