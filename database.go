package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
)

func NewConnection() (*pgx.Conn, error) {
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//dbName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")
	//connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		"shippy.creeozqpgkzm.us-east-1.rds.amazonaws.com", "postgres", "shippy", "123Farit")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
