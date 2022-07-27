package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/x0y14/jackal/database"
	"github.com/x0y14/jackal/gen/notify/v1/notifyv1connect"
	"github.com/x0y14/jackal/mem"
	"github.com/x0y14/jackal/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	err := database.Init(os.Getenv("SQLITE_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	mem.Init(os.Getenv("REDIS_URL"))

	conn, err := amqp.Dial(os.Getenv("RABBIT_URL"))
	if err != nil {
		log.Fatalf("failed to connect rabbitmq: %v", err)
	}
	defer conn.Close()

	notifyServiceHandler := &service.NotifyService{
		Rb: conn,
	}

	mux := http.NewServeMux()
	mux.Handle(notifyv1connect.NewNotifyServiceHandler(
		notifyServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Service listening on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
