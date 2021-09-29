package v1

import (
	"fmt"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr())
	//client, err := api.NewClient(&api.Config{
	//	Address: "db.kylan.pro:8500",
	//})
	//dis := consul.New(client)
	//c, err := rpcx.Dial(
	//	rpcx.WithTimeout(3*time.Second),
	//	rpcx.WithServerName(ServiceNameOfGreeter),
	//	rpcx.WithDiscovery(dis),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//client1 := NewGreeterClient(c)
	//time.Sleep(1 * time.Second)
	//for {
	//	time.Sleep(time.Second)
	//	hello, err := client1.SayHello(context.Background(), &HelloRequest{Name: "123123"})
	//	if err != nil {
	//		panic(err)
	//		return
	//	}
	//	fmt.Println(hello.Message)
	//}
}
