package db

import (
	"database/sql"
	"fmt"
	"log"

	configs "github.com/4lerman/e_com/common/config"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
}

func NewPSQLStorage(cfg *DbConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error occured when connecting to db:", err)
	}

	return db, nil
}

func InitStorage() (*sql.DB, error) {
	db, err := NewPSQLStorage(&DbConfig{
		Host:     configs.Envs.DBAddress,
		User:     configs.Envs.DBUser,
		Port:     configs.Envs.DBPort,
		Dbname:   configs.Envs.DBName,
		Password: configs.Envs.DBPassword,
	})

	if err != nil {
		return nil, fmt.Errorf("DB init error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("DB connection error: %v", err)
	}

	return db, nil
}
