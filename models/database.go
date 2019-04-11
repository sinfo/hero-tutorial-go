package models

import (
	"fmt"

	"github.com/globalsign/mgo"
)

var (
	DB *mgo.Database
)

func InitDB(url string, name string) *mgo.Database {
	var session *mgo.Session
	var err error

	if session, err = mgo.Dial(url); err != nil {
		panic(fmt.Errorf("URL: %v, name: %v, err: %s", url, name, err))
	}

	DB = session.DB(name)
	InitHeroCollection()

	return DB
}
