package server

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kykira/kratos-layout/api/helloworld/v1"
	"github.com/kykira/kratos-layout/internal/conf"
	"github.com/kykira/kratos-layout/internal/service"
	"github.com/kykira/kratos-rpcx-transport/rpcx"
)

// NewRPCXServer new a gRPC server.
func NewRPCXServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *rpcx.Server {
	opts := []rpcx.ServerOption{
		rpcx.Logger(logger),
	}
	if c.Rpcx.Network != "" {
		opts = append(opts, rpcx.Network(c.Rpcx.Network))
	}
	if c.Rpcx.Addr != "" {
		opts = append(opts, rpcx.Address(c.Rpcx.Addr))
	}
	if c.Rpcx.Timeout != nil {
		opts = append(opts, rpcx.Timeout(c.Rpcx.Timeout.AsDuration()))
	}
	//opts = append(opts, rpcx.Plugin())
	srv := rpcx.NewServer(opts...)
	err := v1.RegisterGreeterServe(srv.Server, greeter, "")
	if err != nil {
		panic(err)
		return nil
	}
	return srv
}
