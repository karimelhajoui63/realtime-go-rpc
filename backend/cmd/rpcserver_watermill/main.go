package main

import (
	"log"
	"net/http"

	rpchandler "rpc-server/internal/adapter/inbound/painterusecase"
	"rpc-server/internal/adapter/outbound/broadcasterrepo"
	"rpc-server/internal/adapter/outbound/colorrepo"
	"rpc-server/internal/core/usecase"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

// "rabbitmq" here is the name of the service in the "docker-compose.yml"
// var amqpURI = "amqp://guest:guest@rabbitmq:5672"

// This is the URL from dev container if the rabbitmq container is ran from it
// TODO: manage to adapt the URL from ENV variables
var amqpURI = "amqp://guest:guest@localhost:5673"

func main() {
	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	broadcasterRepo, err := broadcasterrepo.NewWatermillBroadcasterRepository(&amqpConfig, "queue-broadcast")
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

	err = http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}
