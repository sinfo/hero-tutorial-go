package models

import (
	"fmt"

	"github.com/globalsign/mgo"
)

/*
InitDB initializes a DB client instance for the server.
MongoDB database must be accessible by the URL given.
*/
func InitDB(url string, name string) *mgo.Database {
	var session *mgo.Session
	var err error

	if session, err = mgo.Dial(url); err != nil {
		panic(fmt.Errorf("URL: %v, name: %v, err: %s", url, name, err))
	}

	db := session.DB(name)
	InitHeroCollection(db)

	return db
}
