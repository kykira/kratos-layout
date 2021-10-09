package v1

import (
	"context"
	"fmt"
	"github.com/kykira/kratos-rpcx-transport/rpcx"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	c, err := rpcx.Dial(
		rpcx.WithTimeout(3*time.Second),
		rpcx.WithServerName(ServiceNameOfGreeter),
		rpcx.WithEndpoint("0.0.0.0:9000"),
	)
	if err != nil {
		panic(err)
	}
	client1 := NewGreeterClient(c)
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		hello, err := client1.SayHello(context.Background(), &HelloRequest{Name: "123123"})
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(hello.Message)
	}
}
