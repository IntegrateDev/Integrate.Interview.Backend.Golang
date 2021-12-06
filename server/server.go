package server

import (
	"log"
	"net/http"
	"time"
)

func Start(ip, port string) {
	swaggerip := "127.0.0.1"
	swaggerport := "1323"
	srvdocs := &http.Server{
		Handler:      routerdocs,
		Addr:         swaggerip + ":" + swaggerport,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		log.Println(srvdocs.ListenAndServe())
	}()
	log.Printf("Running swagger on %s:%s", swaggerip, swaggerport)

	log.Printf("Running on %s:%s", ip, port)

	srv := &http.Server{
		Handler:      router,
		Addr:         ip + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
