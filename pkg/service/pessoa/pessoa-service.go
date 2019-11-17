package pessoa

import "database/sql"

// versão da api
const (
	apiVersion = "v1"
)

// PessoaServiceSerce implementação de v1.PessoaServiceSerce interface proto
type PessoaServiceSerce struct {
	db *sql.DB
}

// NewPessoaServiceServer Cria o servidor para pessoa
func NewPessoaServiceServer(db *sql.DB) pessoa.PessoaServiceSerce {
	return &pessoaServiceSerce{db: db}
}
