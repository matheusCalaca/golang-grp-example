package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"github.com/matheusCalaca/golanggrpexample/pkg/util"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)

	//Pessoa cliente
	clientePessoa(conn, reminder, err, ctx)
	//Endereco cliente
	clienteEndereco(conn, err, ctx)

}

func clienteEndereco(conn *grpc.ClientConn, err error, ctx context.Context) {
	// Endereço
	clientEndereco := pessoa.NewEnderecoServiceClient(conn)
	reqEndereco := pessoa.CriarEnderecoRequest{
		Api: apiVersion,
		Endereco: &pessoa.Endereco{
			Cep:         74413140,
			Logradouro:  "Rua marechal lino de morais",
			Complemento: "qd 145",
			Bairro:      "cidade jardim",
			Cidade:      "Goiania",
			Uf:          "GO",
		},
	}
	fmt.Println(reqEndereco)
	responseEndereco, err := clientEndereco.CriarEndereco(ctx, &reqEndereco)
	if err != nil {
		log.Fatalf("falha ao criar enderreço %v", err)
	}
	log.Printf("Endereço criado <%+v>\n\n", responseEndereco)
}

func clientePessoa(conn *grpc.ClientConn, reminder *timestamp.Timestamp, err error, ctx context.Context) {

	// cria uma nova conecxão
	c := pessoa.NewPessoaServiceClient(conn)
	// format RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	//pfx := t.Format(time.RFC3339Nano)
	dataNascimento, err := util.DataBRtoProtoBuffDate("27-08-1995")
	if err != nil {
		panic(err)
	}
	// criar uma pessoa
	req1 := pessoa.CriarPessoaRequest{
		Api: apiVersion,
		Pessoa: &pessoa.Pessoa{
			Nome:         "Matheus Calaça 2",
			DtNascimento: dataNascimento,
			Email:        "matheusfcalaca@gmail.com",
			Reminder:     reminder,
		},
	}
	fmt.Println(req1)
	res1, err := c.Criar(ctx, &req1)
	if err != nil {
		log.Fatalf("Falha ao criar uma pessoa: %v", err)
	}
	log.Printf("Pessoa Criada: <%+v>\n\n", res1)
}
