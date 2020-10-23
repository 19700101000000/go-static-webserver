package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "hosting port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("hosting port", *port)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir("./dist"))))
}
