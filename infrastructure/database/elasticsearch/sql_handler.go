package database

import (
	"app/domain"
	"app/interface/config"
	"app/interface/database"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"

	_ "github.com/go-sql-driver/mysql"
)

// SqlHandler is an struct of sql connection
type SqlHandler struct {
	Client *elastic.Client
}

func NewSqlHandler(configHandler config.ConfigHandler) (database.SqlHandler, error) {
	condsn := "http://" + configHandler.GetDatabaseHost() + ":" + configHandler.GetDatabasePort()
	client, err := elastic.NewClient(
		elastic.SetURL(condsn),
	)
	if err != nil {
		return nil, fmt.Errorf("Elastic Client Failed ", err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Client = client
	return sqlHandler, nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	return nil, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	query := elastic.NewTermQuery(statement, args[0].(string))
	rows, err := handler.Client.Search().
		Index("test").
		Type("test").
		Query(query).
		Pretty(true).
		Do(context.Background())

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
	Rows *elastic.SearchResult
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	object := dest[0].(*domain.Test)
	for _, hit := range r.Rows.Hits.Hits {
		json.Unmarshal(*hit.Source, object)
	}
	return nil
}

func (r SqlRow) Next() bool {
	return true
}

func (r SqlRow) Close() error {
	return nil
}
