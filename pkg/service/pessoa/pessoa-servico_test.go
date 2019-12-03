package pessoa

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	pessoa "github.com/matheusCalaca/golanggrpexample/pkg/api/pessoa"
	"reflect"
	"time"

	"testing"
)

func Test_userServiceServer_Create(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := pessoa.NewPessoaServiceServer(db)
	fmt.Printf(s)

	tm := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(tm)

	type args struct {
		ctx context.Context
		req *pessoa.CrearPessoaRequest
	}
	tests := []struct {
		name    string
		s       pessoa.PessoaServiceServer
		args    args
		mock    func()
		want    *pessoa.CriarPessoaResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &pessoa.CrearPessoaRequest{
					Api: "pessoa",
					Pessoa: &pessoa.Pessoa{
						Nome:     "nome",
						Email:    "nome@gmail.com",
						Reminder: reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO pessoa").WithArgs("nome", "email", tm).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &pessoa.CriarPessoaResponse{
				Api: "pessoa",
				Id:  1,
			},
		},
		{
			name: "Unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &pessoa.CrearPessoaRequest{
					Api: "pessoa000",
					Pessoa: &pessoa.Pessoa{
						Nome:  "nome",
						Email: "nome@gmail.com",
						Reminder: &timestamp.Timestamp{
							Seconds: 1,
							Nanos:   -1,
						},
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "Invalid Reminder field format",
			s:    s,
			args: args{
				ctx: ctx,
				req: &pessoa.CrearPessoaRequest{
					Api: "pessoa",
					Pessoa: &pessoa.Pessoa{
						Nome:  "nome",
						Email: "sobrenome",
						Reminder: &timestamp.Timestamp{
							Seconds: 1,
							Nanos:   -1,
						},
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
		{
			name: "INSERT failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &pessoa.CrearPessoaRequest{
					Api: "pessoa",
					Pessoa: &pessoa.Pessoa{
						Nome:     "nome",
						Email:    "sobrenome",
						Reminder: reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO Pessoa").WithArgs("nome", "sobrenome", tm).
					WillReturnError(errors.New("INSERT failed"))
			},
			wantErr: true,
		},
		{
			name: "LastInsertId failed",
			s:    s,
			args: args{
				ctx: ctx,
				req: &pessoa.CrearPessoaRequest{
					Api: "pessoa",
					Pessoa: &pessoa.Pessoa{
						Nome:     "nome",
						Email:    "sobrenome",
						Reminder: reminder,
					},
				},
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO Pessoa").WithArgs("nome", "sobrenome", tm).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("LastInsertId failed")))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Criar(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userServiceServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServiceServer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
