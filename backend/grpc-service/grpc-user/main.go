package main

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:26
//
import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-user/handler"
	"grpc-user/proto"
	"grpc-user/svc"
	"net"
)

var (
	IP         *string
	Port       *int
	configPath string
)

func main() {
	serSvc := svc.NewSrvCtx(configPath)
	fmt.Println(serSvc.Config.MySQl.Datasource)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	reflection.Register(server) // 方便grpcUI 调试
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

func init() {
	IP = flag.String("ip", "0.0.0.0", "IP地址")
	Port = flag.Int("port", 8023, "端口号")
	flag.StringVar(&configPath, "cfg", "dev.yaml", "配置文件路径")
	flag.Parse()
	fmt.Println(*IP, *Port)
}
