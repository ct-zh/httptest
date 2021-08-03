package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	_ = os.Mkdir("/tmp/file", 0755)
	http.Handle("/get/", http.StripPrefix("/get/", http.FileServer(http.Dir("/tmp/file"))))
	err := http.ListenAndServe(":24505", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
