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
func RunServer(ctx context.Context, port string, PessoaAPI pessoa.PessoaServiceServer, EnderecoApi pessoa.EnderecoServiceServer) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	pessoa.RegisterPessoaServiceServer(server, PessoaAPI)

	pessoa.RegisterEnderecoServiceServer(server, EnderecoApi)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("Desligado  gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("iniciado  gRPC server...")
	return server.Serve(listen)
}
