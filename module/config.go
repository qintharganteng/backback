package module

import (
	"os"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tesdb2024",
}

var MongoConn = atdb.MongoConnect(MongoInfo)