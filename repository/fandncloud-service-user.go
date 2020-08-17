package repository

import (
	"context"
	pb "github.com/farit2000/fandncloud-service-user/proto/fandncloud-service-user"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

type PostgresRepository struct {
	conn *pgx.Conn
}

type Repository interface {
	CreateUser(ctx context.Context, user *pb.User) error
	DeleteUser(ctx context.Context, user *pb.User) error
	Update(ctx context.Context, user *pb.User) error
	GetById(ctx context.Context, id string) (*pb.User, error)
	GetAll(ctx context.Context) ([]*pb.User, error)
}

func (repo *PostgresRepository) CreateUser(ctx context.Context, user *pb.User) error {
	var userId pgtype.UUID
	err := repo.conn.QueryRow(ctx, `
		INSERT INTO users 
			(role_id, first_name, last_name, email, phone, password, is_staff, is_blocked, is_deleted)
    	VALUES (1, 'Test', 'Test', 'test@ya.ru', '+79784756349', '12345', true, false, false) RETURNING id;`).
		Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) DeleteUser(ctx context.Context, user *pb.User) error {

}

func (repo *PostgresRepository) Update(ctx context.Context, user *pb.User) error {

}

func (repo *PostgresRepository) GetById(ctx context.Context, id string) (*pb.User, error) {

}

func (repo *PostgresRepository) GetAll(ctx context.Context) ([]*pb.User, error) {

}