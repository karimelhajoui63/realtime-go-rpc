package main

import (
	"log"
	"net/http"
	rpchandler "rpc-server/internal/adapter/inbound/painterusecase"
	"rpc-server/internal/adapter/outbound/broadcasterrepo"
	"rpc-server/internal/adapter/outbound/colorrepo"
	"rpc-server/internal/core/usecase"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel:", err)
	}
	defer ch.Close()

	broadcasterRepo, err := broadcasterrepo.NewRabbitmqBroadcasterRepository(ch, "queue-broadcast")
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	colorRepo, err := colorrepo.NewInMemoryColorRepository()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	painterUseCase := usecase.NewPainterUseCase(broadcasterRepo, colorRepo)

	mux := http.NewServeMux()
	handler := rpchandler.NewRpcPainterHandler(mux, painterUseCase)

	err = http.ListenAndServe("localhost:8080", handler)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}
