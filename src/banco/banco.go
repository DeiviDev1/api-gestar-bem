package banco

import (
	"api-gestar-bem/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver do MySQL
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectBD)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
