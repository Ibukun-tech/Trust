package main

import (
	"net/http"

	db "github.con/Ibukun-tech/trust/Db"
)

func main() {
	db.NewDbConnection()

	s := http.Server{}
	s.ListenAndServe()
}
