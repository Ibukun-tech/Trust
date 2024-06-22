package db

import (
	"database/sql"

	"github.con/Ibukun-tech/trust/utils"
)

type DbConnect struct {
	Db *sql.DB
}

// I am must have a connection that must like take in everything so that each associated handlers can make use of it a nd all of it finish

func NewDbConnection(pd string) (*DbConnect, error) {
	connect := utils.ConnectDb()
	db, err := sql.Open(pd, connect)
	if err != nil {
		return nil, err
	}
	return &DbConnect{
		Db: db,
	}, nil
}
