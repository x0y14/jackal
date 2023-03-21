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
	"time"
)

func main() {
	err := database.Init(os.Getenv("SQLITE_PATH"))
	if err != nil {
		log.Fatalf("failed to connect to sqlite: %s", err)
	}
	defer database.Close()

	// REDISとの接続
	redisTryCount := 0
	for redisTryCount < 10 {
		redisTryCount++
		if err = mem.Init(os.Getenv("REDIS_URL")); err != nil {
			time.Sleep(time.Second * 10)
			continue
		} else {
			redisTryCount = -1
			break
		}
	}
	if redisTryCount != -1 {
		log.Fatalf("failed to connect to redis in 10 times: %s", err)
	}

	// RABBITMQとの接続
	rabbitTryCount := 0
	var conn *amqp.Connection
	for rabbitTryCount < 10 {
		rabbitTryCount++
		conn, err = amqp.Dial(os.Getenv("RABBIT_URL"))
		if err != nil {
			time.Sleep(time.Second * 10)
			continue
		} else {
			rabbitTryCount = -1
			break
		}
	}
	if rabbitTryCount != -1 {
		log.Fatalf("failed to connect to rabbitMq in 10 times: %s", err)
	}
	defer conn.Close()

	log.Printf("connect to redis & rabbitMq successfuly")

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
