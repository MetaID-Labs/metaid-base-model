package major

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DsMetaIdDatabaseRW = "metanetRW"
)

func GetMetaIdReadDB() (*mongo.Client, error) {
	return GetDBWith(DsMetaIdDatabaseRW)
}

func GetMetaIdReadWriteDB() (*mongo.Client, error) {
	return GetDBWith(DsMetaIdDatabaseRW)
}
