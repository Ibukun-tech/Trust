package utils

import (
	"fmt"

	model "github.con/Ibukun-tech/trust/Internals/Models"
)

func ConnectDb(c model.ConfigDatabase) string {
	// return postgress://root:secret@localhost:5432/trustDb?
	return fmt.Sprintf("postgress://%s:%s@%s:%s/%s?sslmode= disable", c.DbUser, c.DbPass, c.DbRemote, c.DbPort, c.DbName)
}
