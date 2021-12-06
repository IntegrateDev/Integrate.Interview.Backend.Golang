package server

import (
	"log"
	"net/http"
	"time"
)

func Start(ip, port string) {
	srvdocs := &http.Server{
		Handler:      routerdocs,
		Addr:         "127.0.0.1:1323",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		log.Println(srvdocs.ListenAndServe())
	}()

	srv := &http.Server{
		Handler:      router,
		Addr:         ip + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
