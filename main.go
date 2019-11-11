package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/microservicedemo/echo/internal/impl/service"
	pb "github.com/microservicedemo/echo/service"
	"google.golang.org/grpc"
)

func main() {
	addr := ":9090"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &service.EchoImpl{})

	log.Println("starting...", addr)
	s.Serve(l)
}

func httpHandle() {
	gmux := runtime.NewServeMux()

	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gmux, "echo:9090", opts)
	if err != nil {
		log.Fatal(err)
		return
	}
	gmux.GetForwardResponseOptions()
	httpmux := http.NewServeMux()
	httpmux.Handle("/", gmux)
	http.ListenAndServe(":9091", httpmux)
}
