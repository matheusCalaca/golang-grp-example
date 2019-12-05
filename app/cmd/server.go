package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheusCalaca/golanggrpexample/app/cmd/grpc"
	"github.com/matheusCalaca/golanggrpexample/app/interface/rpc/service/pessoa"
)

// Config configuração para o servidor
type Config struct {

	// GRPCPort porta que o grpc vai ficar escutndo
	GRPCPort string

	// DB Datastore parameters
	// DatastoreDBHost host do BD
	DatastoreDBHost string
	// DatastoreDBUser username do BD
	DatastoreDBUser string
	// DatastoreDBPassword senha do BD
	DatastoreDBPassword string
	// DatastoreDBSchema schema do BD
	DatastoreDBSchema string
}

// RunServer gerenciar server GRPC e HTTP gatwarey
func RunServer() error {
	ctx := context.Background()

	// Obtem a configuração do bd
	// todo: trocar por properties
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to start")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("Porta invalisa para o gRPC: '%s'", cfg.GRPCPort)
	}

	// Especifica o parametros da data para o MYSQL
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	PessoaAPI := pessoa.NewPessoaServiceServer(db)
	EnderecoAPI := pessoa.NewEnderecoServiceServer(db)

	return grpc.RunServer(ctx, cfg.GRPCPort, PessoaAPI, EnderecoAPI)
}
