package main

import (
	"flag"
	"fmt"
	"ordering-platform/pkg/interceptor/rpcserver"
	"ordering-platform/rpc/sys/internal/config"
	apiServiceServer "ordering-platform/rpc/sys/internal/server/apiservice"
	deptServiceServer "ordering-platform/rpc/sys/internal/server/deptservice"
	menuServiceServer "ordering-platform/rpc/sys/internal/server/menuservice"
	roleServiceServer "ordering-platform/rpc/sys/internal/server/roleservice"
	tokenServiceServer "ordering-platform/rpc/sys/internal/server/tokenservice"
	userServiceServer "ordering-platform/rpc/sys/internal/server/userservice"
	"ordering-platform/rpc/sys/internal/svc"
	"ordering-platform/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sysclient.RegisterUserServiceServer(grpcServer, userServiceServer.NewUserServiceServer(ctx))
		sysclient.RegisterApiServiceServer(grpcServer, apiServiceServer.NewApiServiceServer(ctx))
		sysclient.RegisterMenuServiceServer(grpcServer, menuServiceServer.NewMenuServiceServer(ctx))
		sysclient.RegisterDeptServiceServer(grpcServer, deptServiceServer.NewDeptServiceServer(ctx))
		sysclient.RegisterRoleServiceServer(grpcServer, roleServiceServer.NewRoleServiceServer(ctx))
		sysclient.RegisterTokenServiceServer(grpcServer, tokenServiceServer.NewTokenServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpcserver.LogInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
