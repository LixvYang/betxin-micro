package main

import (
	"flag"
	"fmt"

	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/config"
	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/server"
	"github.com/lixvyang/betxin-micro/service/category/rpc/internal/svc"
	"github.com/lixvyang/betxin-micro/service/category/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/categorysrv.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCategorysrvServer(grpcServer, server.NewCategorysrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 将服务注册给consul
	if err := consul.RegisterService(c.ListenOn, c.Consul); err != nil {
		panic(err)
	}
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
