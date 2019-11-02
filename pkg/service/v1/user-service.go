package v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/matheusCalaca/golang-grp-example/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// userServiceServer is implementation of v1.UserServiceServer proto interface
type userServiceServer struct {
	db *sql.DB
}

// NewUserServiceServer creates User service
func NewUserServiceServer(db *sql.DB) v1.UserServiceServer {
	return &userServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *userServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *userServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new todo task
func (s *userServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.User.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// insert User entity data
	res, err := c.ExecContext(ctx, "INSERT INTO User(`nome`, `sobrenome`, `Reminder`) VALUES(?, ?, ?)",
		req.User.nome, req.User.sobrenome, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into User-> "+err.Error())
	}

	// get ID of creates User
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created User-> "+err.Error())
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Read todo task
func (s *userServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// query User by ID
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `nome`, `sobrenome`, `Reminder` FROM User WHERE `ID`=?",
		req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
			req.Id))
	}

	// get User data
	var td v1.User
	var reminder time.Time
	if err := rows.Scan(&td.Id, &td.nome, &td.sobrenome, &reminder); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
	}
	td.Reminder, err = ptypes.TimestampProto(reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple User rows with ID='%d'",
			req.Id))
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		User: &td,
	}, nil

}

// Update todo task
func (s *userServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.User.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// update User
	res, err := c.ExecContext(ctx, "UPDATE User SET `nome`=?, `sobrenome`=?, `Reminder`=? WHERE `ID`=?",
		req.User.nome, req.User.sobrenome, reminder, req.User.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update User-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
			req.User.Id))
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete todo task
func (s *userServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// delete User
	res, err := c.ExecContext(ctx, "DELETE FROM User WHERE `ID`=?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete User-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with ID='%d' is not found",
			req.Id))
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}

// Read all todo tasks
func (s *userServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get User list
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `nome`, `sobrenome`, `Reminder` FROM User")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from User-> "+err.Error())
	}
	defer rows.Close()

	var reminder time.Time
	list := []*v1.User{}
	for rows.Next() {
		td := new(v1.User)
		if err := rows.Scan(&td.Id, &td.nome, &td.sobrenome, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from User row-> "+err.Error())
		}
		td.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from User-> "+err.Error())
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		Users: list,
	}, nil
}