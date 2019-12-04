package pessoa

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// versão da api
const (
	apiVersion = "pessoa"
)

// pessoaServiceService implementação de pessoa.PessoaServiceServer interface proto
type pessoaServiceService struct {
	db *sql.DB
}

//// NewPessoaServiceServer Cria o servidor para pessoa
func NewPessoaServiceServer(db *sql.DB) pessoa.PessoaServiceServer {
	return &pessoaServiceService{db: db}
}

// checkAPI verifica se a versão da api do cliente e suportada pelo o servidor
func (s *pessoaServiceService) checkAPI(api string) error {

	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect retorna o pool de conexao com o database
func (s *pessoaServiceService) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create nova pessoa
func (s *pessoaServiceService) Criar(ctx context.Context, req *pessoa.CriarPessoaRequest) (*pessoa.CriarPessoaResponse, error) {
	// check verifica se a versão da api do cliente e suportada pelo o servidor
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// obtem a conexao com o BD
	con, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer con.Close()

	// start conexão com o server
	connGrpc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("não foi possivel conectar : %v", err)
	}
	defer connGrpc.Close()

	// adicionando endereço para a pessoa
	pessoa.NewEnderecoServiceClient(connGrpc)

	//Endereco cliente
	enderecoResponse, err := clienteEndereco(connGrpc, err, ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	// set a data atual ao redmier
	reminder, err := ptypes.Timestamp(req.Pessoa.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder campo com o formato invalido -> "+err.Error())
	}

	dtNascimento, err := ptypes.Timestamp(req.Pessoa.DtNascimento)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Data de Nascimento Formato invalido -> "+err.Error())
	}

	// insert no bd os dados da pessoa
	res, err := con.ExecContext(ctx, "INSERT INTO pessoa (`NOME`, `DATA_NASCIMENTO`, `EMAIL`, `ENDERECO_ID`,`REMIDER`)	VALUES	(?,	?,?,  ?, ? )",
		req.Pessoa.Nome, dtNascimento, req.Pessoa.Email, enderecoResponse.Id, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Falha ao inserir pessoa no banco de dados-> "+err.Error())
	}

	// obtem o ultimo ai inserido
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "falha ao trazer o ultimo id-> "+err.Error())
	}

	return &pessoa.CriarPessoaResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func clienteEndereco(conn *grpc.ClientConn, err error, ctx context.Context) (*pessoa.CriarEnderecoResponse, error) {
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
		return nil, status.Error(codes.Unknown, "Erro ao criar Endereço ->  "+err.Error())
	}
	log.Printf("Endereço criado <%+v>\n\n", responseEndereco)

	return responseEndereco, nil
}
