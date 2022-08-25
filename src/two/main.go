package main

import (
	"net/http"

	"github.com/fox-wei/gowebframe/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core, //?Handle 是接口
		Addr:    ":9090",
	}

	server.ListenAndServe()
}
