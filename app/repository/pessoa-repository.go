package repository

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	"github.com/matheusCalaca/golanggrpexample/app/interface/rpc/api/pessoa"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// pessoaServiceService implementação de pessoa.PessoaServiceServer interface proto
type pessoaServiceService struct {
	db *sql.DB
}

const (
	apiVersion = "pessoa"
)

//// connect retorna o pool de conexao com o database
//func (service *pessoaServiceService) connect(ctx context.Context) (*sql.Conn, error) {
//	c, err := service.db.Conn(ctx)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
//	}
//	return c, nil
//}

/**
* CriarIdentificadorRepository insere o identificador no banco de dados
 */
func CriarIdentificadorRepository(conn *sql.Conn, identificador *pessoa.Identificador, ctx context.Context) (*pessoa.CriarIdentificadorResponse, error) {
	// obtem a conexao com o BD
	_, err := conn.ExecContext(ctx, "INSERT INTO identificador (`CPF`, `RG`) values (?,?)", identificador.Cpf, identificador.Rg)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Erro ao inserir Identificador da pessoa no banco de dados -> "+err.Error())
	}

	response := &pessoa.CriarIdentificadorResponse{Cpf: identificador.Cpf, Api: apiVersion}
	return response, nil
}

/**
* CriarTelefoneRepository insere o telefone no banco de dados
 */
func CriarTelefoneRepository(conn *sql.Conn, telefone *pessoa.Telefone, ctx context.Context) (*pessoa.CriarTelefoneResponse, error) {
	result, err := conn.ExecContext(ctx, "insert into telefone (TIPO, DD, NUMERO) values (?,?,?)", telefone.Tipo, telefone.Dd, telefone.Numero)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Erro ao inserir o Telefone -> "+err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Erro ao buscar o ultimo ID inserido no Telefone -> "+err.Error())
	}

	response := &pessoa.CriarTelefoneResponse{
		Api: apiVersion,
		Id:  id,
	}
	return response, nil
}

func CriarPessoaRepository(con *sql.Conn, ctx context.Context, reqPessoa *pessoa.Pessoa, enderecoResponse *pessoa.CriarEnderecoResponse, identificadorResponse *pessoa.CriarIdentificadorResponse, telefoneResponse *pessoa.CriarTelefoneResponse) (*pessoa.CriarPessoaResponse, error) {
	// set a data atual ao redmier
	reminder, err := ptypes.Timestamp(reqPessoa.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder campo com o formato invalido -> "+err.Error())
	}
	dtNascimento, err := ptypes.Timestamp(reqPessoa.DtNascimento)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Data de Nascimento Formato invalido -> "+err.Error())
	}
	// insert no bd os dados da pessoa
	res, err := con.ExecContext(ctx, "INSERT INTO pessoa (`NOME`, `DATA_NASCIMENTO`, `EMAIL`, `ENDERECO_ID`,`IEDNTIFICADOR_ID`,`TELEFONE_ID`,`REMIDER`)	VALUES	(?, ?, ?,?,?, ?, ? )",
		reqPessoa.Nome, dtNascimento, reqPessoa.Email, enderecoResponse.Id, identificadorResponse.Cpf, telefoneResponse.Id, reminder)
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
