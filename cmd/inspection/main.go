package main

import (
	"flag"
	"log"
	"net/http"
	"warehouse5s/internal/api"
	"warehouse5s/internal/service"
	"warehouse5s/internal/storage"
)

func main() {
	path := flag.String("db", "inspection.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, e := storage.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := &service.Service{Store: st, Clock: service.FixedClock{Value: "2026-01-01T00:00:00Z"}}
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, api.Server{Svc: svc}.Routes()))
}
