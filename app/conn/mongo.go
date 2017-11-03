package conn

import (
	"github.com/go-bongo/bongo"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var conn *bongo.Connection
var err error

func GetConnection() *bongo.Connection {
	if conn == nil {
		conn, err = bongo.Connect(&bongo.Config{
			ConnectionString: viper.GetString("databases.mongodb.uri"),
			Database:         viper.GetString("databases.mongodb.dbname"),
		})
		if err != nil {
			panic(err)
		}
		return conn
	}
	return conn
}

func GetCollection(name string) *bongo.Collection {
	return GetConnection().Collection(name)
}

func GetMGoCollection(name string) *mgo.Collection {
	return GetConnection().Session.DB(viper.GetString("databases.mongodb.dbname")).C(name)
}
