package main

import (
	"api-grpc/api/resthandlers"
	"api-grpc/api/routes"
	"api-grpc/pb"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	port     int
	authaddr string
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authaddr, "authaddr", "localhost:9001", "authentication service address")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(authaddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	authSvcClient := pb.NewAuthServiceClient(conn)
	authHandlers := resthandlers.NewAuthHandlers(authSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("API service running on [::]:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
