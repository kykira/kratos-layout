package v1

import (
	"context"
	"fmt"
	consul "github.com/go-kratos/consul/registry"
	"github.com/hashicorp/consul/api"
	"github.com/kykira/kratos-rpcx-transport/rpcx"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client, err := api.NewClient(&api.Config{
		Address: "db.kylan.pro:8500",
	})
	dis := consul.New(client)
	c, err := rpcx.Dial(
		rpcx.WithTimeout(3*time.Second),
		rpcx.WithServerName(ServiceNameOfGreeter),
		rpcx.WithEndpoint("0.0.0.0:9000"),
		rpcx.WithDiscovery(dis),
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
