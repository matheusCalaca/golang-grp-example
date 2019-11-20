package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"google.golang.org/grpc"
)

// RunServer inicia um servidor gRPC
func RunServer(ctx context.Context, PessoaAPI pessoa.PessoaServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	pessoa.RegisterPessoaServiceServer(server, PessoaAPI)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
