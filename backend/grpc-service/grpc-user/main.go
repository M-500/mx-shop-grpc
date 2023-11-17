package main

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:26
//
import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"grpc-user/handler"
	"grpc-user/pkg/general"
	"grpc-user/pkg/utils/str"
	"grpc-user/proto"
	"grpc-user/svc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	IP         *string
	Port       *int
	configPath string
)

func main() {
	serSvc := svc.NewSrvCtx(configPath)
	cfg := serSvc.Config
	srvID := str.UUID()
	consulAddr := fmt.Sprintf("%s:%d", cfg.ConsulCfg.Host, cfg.ConsulCfg.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	reflection.Register(server) // 方便grpcUI 调试

	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	// 注册服务
	general.Register(*IP, *Port, cfg.Name, []string{"test", "fuck"}, srvID, consulAddr)

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = general.DeregisterService(srvID, consulAddr); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}

func init() {
	IP = flag.String("ip", "192.168.1.51", "IP地址")
	Port = flag.Int("port", 50051, "端口号")
	flag.StringVar(&configPath, "cfg", "dev.yaml", "配置文件路径")
	flag.Parse()
	fmt.Println(*IP, *Port)
}
