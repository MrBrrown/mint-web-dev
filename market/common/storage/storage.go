package rqstorage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbInfo struct {
	Host  string `json:"host" yaml:"host"`
	Port  string `json:"port" yaml:"port"`
	Login string `json:"login" yaml:"login"`
	Pass  string `json:"pass" yaml:"pass"`
	Db    string `json:"db" yaml:"db"`
}

type DataBase struct {
	DB *sql.DB
}

func New(info DbInfo) (*DataBase, error) {
	db_, err := connectToDB(info)
	if err != nil {
		return nil, err
	}

	return &DataBase{DB: db_}, nil
}

func (d *DataBase) Close() {
	d.DB.Close()
}

func connectToDB(info DbInfo) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", info.Login, info.Pass, info.Host, info.Port, info.Db)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
