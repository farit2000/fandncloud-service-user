package repository

import (
	"context"
	pb "github.com/farit2000/FandNCloud/fandncloud-service-user/proto/fandncloud-service-user"
	"github.com/jackc/pgx"
)

type PostgresRepository struct {
	conn *pgx.Conn
}

type Repository interface {
	Create(ctx context.Context, user *pb.User) error
	Delete(ctx context.Context, user *pb.User) error
	Block(ctx context.Context, user *pb.User) error
	UnBlock(ctx context.Context, user *pb.User) error
	Update(ctx context.Context, user *pb.User) error
	GetById(ctx context.Context, id string) (*pb.User, error)
	GetByEmail(ctx context.Context, email string) (*pb.User, error)
	GetByPhone(ctx context.Context, email string) (*pb.User, error)
	GetAll(ctx context.Context) ([]*pb.User, error)
}