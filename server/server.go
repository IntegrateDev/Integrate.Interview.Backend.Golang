package server

import (
	"log"
	"net/http"
	"time"

	"Integrate.Interview.Backend.Golang/server/controllers"
)

func Start(ip, port string, service interface{}) {
	swaggerip := "127.0.0.1"
	swaggerport := "1323"

	controllersManager := controllers.NewController(&controllers.ControllerManager{Service: service})
	apirouter, docrouter := Api(controllersManager)

	srvdocs := &http.Server{
		Handler:      docrouter,
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
		Handler:      apirouter,
		Addr:         ip + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
