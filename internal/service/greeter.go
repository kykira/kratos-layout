package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kykira/kratos-layout/api/helloworld/v1"
	"github.com/kykira/kratos-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	Name string
	uc   *biz.GreeterUsecase
	log  *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{Name: v1.ServiceNameOfGreeter, uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest, reply *v1.HelloReply) error {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "error" {
		return errors.New(400, "123123", "123")
	}
	*reply = v1.HelloReply{Message: "Hello " + in.GetName()}
	return nil
}
