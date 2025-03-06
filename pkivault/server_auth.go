package pkivault

import (
	"context"
	"database/sql"
	fmt "fmt"

	grpctoken "github.com/ZolaraProject/library/grpctoken"
	logger "github.com/ZolaraProject/library/logger"
	. "github.com/ZolaraProject/pki-vault-service/pkivaultrpc"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/metadata"
)

func (*server) CreateUser(ctx context.Context, req *UserCreateRequest) (*Response, error) {
	grpcToken := grpctoken.GetToken(ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Err(grpcToken, "metadata not found")
	}
	logger.Debug(grpcToken, "metadata: %+v", md)

	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err(grpcToken, "Open error : %v", err)
		return nil, err
	}

	// if req.Username == "" || req.Email == "" || (req.Password == "" && !req.IsOAuth) {
	// 	logger.Err(grpcToken, "all the fields are mandatory")
	// 	return nil, fmt.Errorf("all the fields are mandatory")
	// }

	if req.Username == "" || req.Email == "" {
		logger.Err(grpcToken, "all the fields are mandatory")
		return nil, fmt.Errorf("all the fields are mandatory")
	}

	query := "INSERT INTO users (username, email, password, oauth) VALUES ($1, $2, $3, $4) RETURNING id"
	createUserParams := []interface{}{req.Username, req.Email, req.Password}

	if req.IsOAuth {
		createUserParams = append(createUserParams, true)
	} else {
		createUserParams = append(createUserParams, false)
	}

	var userId sql.NullInt64
	err = db.QueryRow(query, createUserParams...).Scan(&userId)
	if err != nil {
		logger.Err(grpcToken, "failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &Response{
		Message:   "User created successfully",
		CreatedId: userId.Int64,
	}, nil
}

func (*server) DeleteUser(ctx context.Context, req *UserDeleteRequest) (*Response, error) {
	grpcToken := grpctoken.GetToken(ctx)

	db, err := sql.Open("postgres", DbUrl())
	if err != nil {
		logger.Err(grpcToken, "Open error : %v", err)
		return nil, err
	}

	if req.Id == 0 {
		logger.Err(grpcToken, "id is mandatory")
		return nil, fmt.Errorf("id is mandatory")
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err = db.Exec(query, req.Id)
	if err != nil {
		logger.Err(grpcToken, "failed to execute query: %s", err)
		return nil, fmt.Errorf("failed to execute query: %s", err)
	}

	return &Response{
		Message: "User deleted successfully",
	}, nil
}
