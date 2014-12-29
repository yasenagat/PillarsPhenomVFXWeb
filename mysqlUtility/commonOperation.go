package mysqlUtility

import (
	"database/sql"
)

// begin with capitial wordd so it can be accessed by outer
var DBConn *sql.DB

func init() {
	DBConn = ConnectToDB()
}

func ConnectToDB() *sql.DB {
	// connection already exist
	if DBConn != nil {
		return DBConn
	}

	// connection not exist
	prpertyMap := utility
}
