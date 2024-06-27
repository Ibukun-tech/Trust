package main

import (
	"net/http"

	db "github.con/Ibukun-tech/trust/Db"
)

func init() {

	db.NewDbConnection()
}
func main() {

	s := http.Server{}
	s.ListenAndServe()
}
