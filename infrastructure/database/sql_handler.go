package database

import (
	"app/interface/config"
	"app/interface/database"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// SqlHandler is an struct of sql connection
type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(configHandler config.ConfigHandler) (database.SqlHandler, error) {
	condsn := configHandler.GetDatabaseUser() + ":" + configHandler.GetDatabasePassword() + "@tcp(" + configHandler.GetDatabaseHost() + ":" + configHandler.GetDatabasePort() + ")/" + configHandler.GetDatabase()
	conn, err := sql.Open("mysql", condsn)

	if err != nil {
		return nil, fmt.Errorf("DB Connection Configure Error %s", err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler, nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil

}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil

}

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
