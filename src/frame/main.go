package main

import (
	"log"
	"net/http"
)

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(reponse http.ResponseWriter, request *http.Request) {
	// fmt.Fprintln(reponse, "Hello ", request.URL.Path)
	log.Println(request.URL.Path)
}

func main() {
	server := http.Server{
		Addr:    ":8081",   //*监听地址
		Handler: NewCore(), //*核心处理函数
	}

	server.ListenAndServe()
}
