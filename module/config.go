package module

import (
	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tesdb2024",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

