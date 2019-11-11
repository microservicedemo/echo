package main

import (
	"fmt"
	"time"

	echo "github.com/microservicedemo/echo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	cc, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}

	cli := echo.NewEchoClient(cc)

	for i := 0; i < 5; i++ {
		ret, err := cli.Hi(context.Background(), &echo.SimpleMessage{
			MsgId:   "id123",
			Message: "msg43546" + time.Now().String(),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ret)
	}
}
