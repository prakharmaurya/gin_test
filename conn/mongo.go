package conn

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("No DB Server found session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
