package main

//go:generate protoc --proto_path=$GOPATH/src:../proto --micro_out=./proto/ --go_out=./proto/ ../proto/Event.proto

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	_ "github.com/micro/go-plugins/broker/rabbitmq"

	"github.com/pijalu/micro.broker/proto"
)

type sub struct{}

func (s *sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[Sub] Got event %+v with metadata %+v\n", event, md)
	return nil
}

func main() {
	log.Print("Starting subcriber")
	service := micro.NewService(
		micro.Name("go.micro.sub.BrokerTes"))
	service.Init()

	option := func(o *server.SubscriberOptions) {
		o.Queue = "myQueue"
	}

	micro.RegisterSubscriber("topic.events", service.Server(), new(sub), option)

	if err := service.Run(); err != nil {
		log.Panicf("Error: %+v", err)
	}
}
