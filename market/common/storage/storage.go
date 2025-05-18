package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbInfo struct {
	Host  string `json:"host" yaml:"host"`
	Port  string `json:"port" yaml:"port"`
	Login string `json:"login" yaml:"login"`
	Pass  string `json:"pass" yaml:"pass"`
	Db    string `json:"db" yaml:"db"`
	SSL   string `json:"sslmode" yaml:"sslmode"` // "disable", "require", "verify-full"
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
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		info.Host,
		info.Port,
		info.Login,
		info.Pass,
		info.Db,
		info.SSL,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
