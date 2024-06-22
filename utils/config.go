package utils

import "fmt"

func ConnectDb(dbUser, dbPass, dbRemote, dbPort, dbName string) string {
	// return postgress://root:secret@localhost:5432/trustDb?
	return fmt.Sprintf("postgress://%s:%s@%s:%s/%s?sslmode= disable", dbUser, dbPass, dbRemote, dbPort, dbName)
}
