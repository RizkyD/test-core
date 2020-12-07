package artemis

import (
	"database/sql"
)

var ServerAttribute serverAttribute

type serverAttribute struct {
	DBConnection *sql.DB
}

func SetServerAttribute(param string, connectionstr string) {
	ServerAttribute.DBConnection = GetDbConnection(param, connectionstr)
}
