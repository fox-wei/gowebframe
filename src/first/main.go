package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func fooHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, foo")
}

func main() {
	//?创建web程序两个方式
	http.Handle("/foo", http.HandlerFunc(fooHandle))

	http.HandleFunc("/hi", fooHandle)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
