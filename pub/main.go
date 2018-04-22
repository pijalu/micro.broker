package main

//go:generate protoc --proto_path=$GOPATH/src:../proto --micro_out=./proto --go_out=./proto ../proto/Event.proto

import (
	"context"
	"fmt"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/pijalu/micro.broker/proto"
)

func main() {
	log.Printf("Starting Publisher")
	service := micro.NewService(
		micro.Name("go.micro.cli.BrokerTest"))
	service.Init()

	publisher := micro.NewPublisher("topic.events", service.Client())

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		evt := &proto.Event{
			Name: fmt.Sprintf("Hello World %d!", i+1),
		}

		if err := publisher.Publish(ctx, evt); err != nil {
			log.Fatalf("Error: %+v", err)
		} else {
			log.Printf("Sent: %+v", evt)
		}
	}
	select {
	case <-ctx.Done():
		return
	}
}
