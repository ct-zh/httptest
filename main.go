package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		_, _ = writer.Write([]byte(fmt.Sprintf("server[%s]: thank you, %s!", r.Host, r.RemoteAddr)))
		log.Printf("[%s %s]get request from %s", time.Now().Format("2006-01-02 15:04:05"), r.Host, r.RemoteAddr)
	})

	err := http.ListenAndServe(":11234", nil)
	if err != nil {
		panic(err)
	}
}
