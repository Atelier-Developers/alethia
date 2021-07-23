package RPC

import (
	"fmt"
	"github.com/Atelier-Developers/alethia/RPC/impl"
	"github.com/Atelier-Developers/alethia/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

func GetNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

func CreateGrpcServer() {
	portEnv := os.Getenv("GRPC_PORT")
	log.Println("Grpc is Starting...")
	port, err := strconv.ParseUint(portEnv, 10, 64)
	if err != nil {
		log.Fatalf("No grpc port available")
	}
	lis := GetNetListener(uint(port))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	newBirthdayUpdaterServiceGrpcImpl := impl.NewBirthdayUpdaterServiceGrpcImpl()
	proto.RegisterBirthdayUpdaterServer(grpcServer, newBirthdayUpdaterServiceGrpcImpl)
	go func() {
		log.Println("Grpc is Running")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
}
