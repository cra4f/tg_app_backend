package main

import (
	"fmt"
	"log"
	"tg_app_backend/internal/server"
	"tg_app_backend/internal/storage/postgresql"
)

type DbConnectInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func main() {
	//config.MustLoad()
	dbInfo := DbConnectInfo{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "111222",
		Name:     "tg_app",
	}

	pgConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	// как-то по красивее оформить err
	db, err := postgresql.New(pgConnect)

	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(db)
	err = srv.Start()

	if err != nil {
		log.Fatal(err)
	}
}
