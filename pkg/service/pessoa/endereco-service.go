package pessoa

import (
	"database/sql"
	"github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type enderecoServiceService struct {
	db *sql.DB
}

func NewEnderecoServiceServer(db *sql.DB) pessoa.EnderecoServiceServer {
	return &enderecoServiceService{db: db}
}

func (service *enderecoServiceService) connect(ctx context.Context) (*sql.Conn, error) {
	con, err := service.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Falha ao connectar no banco de dados -> "+err.Error())
	}
	return con, nil
}

func (service *enderecoServiceService) CriarEndereco(ctx context.Context, req *pessoa.CriarEnderecoRequest) (*pessoa.CriarEnderecoResponse, error) {

	//abrir conexão
	conn, err := service.connect(ctx)
	if err != nil {
		return nil, err
	}

	result, err := conn.ExecContext(ctx, "insert into endereco(CEP, LOGRADOURO, COMPLEMENTO, BAIRRO, CIDADE, UF) values (?,?,?,?,?,?)",
		req.Endereco.Cep, req.Endereco.Logradouro, req.Endereco.Complemento, req.Endereco.Bairro, req.Endereco.Cidade, req.Endereco.Uf)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Falha ao inserir endereço no banco de dados -> "+err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Falha ao trazer o Ultimo ID -> "+err.Error())
	}

	return &pessoa.CriarEnderecoResponse{
		Api: apiVersion,
		Id:  id,
	}, nil

}
