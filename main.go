package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	addr       = flag.String("addr", ":8080", "http service address")
	indexTempl *template.Template
)

func main() {
	flag.Parse()

	indexTempl = template.Must(template.ParseFiles("index.html"))

	go h.run()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		indexTempl.Execute(w, req.Host)
	})

	http.HandleFunc("/ws", wsHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
