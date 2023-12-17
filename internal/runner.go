package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/icrxz/grpc-go/internal/config"
	"github.com/icrxz/grpc-go/internal/services"
	greeting_v1 "github.com/icrxz/grpc-go/pkg/pb/greeting/v1"
	"google.golang.org/grpc"
)

func Run() error {
	listener := config.InitListener()

	server := grpc.NewServer()
	log.Println("gRPC server initialized")

	greetingService := services.NewGreeterService()
	greeting_v1.RegisterGreeterServiceServer(server, greetingService)
	log.Println("gRPC handlers registered")

	go signalsListener(server)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sigs

	log.Println("Gracefully stopping server...")
	server.GracefulStop()
}
