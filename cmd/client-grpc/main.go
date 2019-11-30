package main

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes"
	v1 "github.com/matheusCalaca/golang-grp-example/pkg/api/v1"
	"github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "pessoa"
)


func main() {
	// obter as configurações
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("não foi possivel conectar : %v", err)
	}
	defer conn.Close()

	// cria uma nova conecxão
	c := pessoa.NewPessoaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)

	// format RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	//pfx := t.Format(time.RFC3339Nano)
	date, _ := time.Parse("2006-01-02", "1995-08-27")

	// crira uma pessoa
	req1 := pessoa.CrearPessoaRequest{
		Api: apiVersion,
		Pessoa: &pessoa.Pessoa{
			Nome:"Matheus Calaça",
			DtNascimento: date.UnixNano(),
			Email: "matheusfcalaca@gmail.com",
			Reminder:    reminder,
		},
	}
	res1, err := c.Criar(ctx, &req1)
	if err != nil {
		log.Fatalf("Falha ao criar uma pessoa: %v", err)
	}
	log.Printf("Pessoa Criada: <%+v>\n\n", res1)


}