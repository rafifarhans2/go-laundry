package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConnection interface {
	Conn() *sql.DB
}

type dbConnection struct {
	db  *sql.DB
	cfg *Config
}

func (d *dbConnection) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.cfg.Dbconfig.Host,
		d.cfg.Dbconfig.Port,
		d.cfg.Dbconfig.User,
		d.cfg.Dbconfig.Password,
		d.cfg.Dbconfig.Name,
	)
	db, err := sql.Open(d.cfg.Dbconfig.Driver, dsn)
	if err != nil {
		println(err)
		return err
	}
	d.db = db
	return nil
}

func (d *dbConnection) Conn() *sql.DB {
	return d.db
}

// Constructor
func NewDbConnection(cfg *Config) (DbConnection, error) {
	conn := &dbConnection{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
