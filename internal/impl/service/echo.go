package service

import (
	"log"

	pb "github.com/microservicedemo/echo/service"
	"golang.org/x/net/context"
)

type EchoImpl struct {
}

func (e *EchoImpl) Echo(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Printf("%+v", in)
	return &pb.SimpleMessage{MsgId: "reply-" + in.MsgId, Message: "hi."}, nil
}
