package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Hero struct {
	_ID  bson.ObjectId `bson:"_id"`
	ID   int           `json:"id" bson:"id"`
	Name string        `json:"name" bson:"name"`
}

var heroCollection *mgo.Collection

func InitHeroCollection() {
	heroCollection = DB.C("heroes")

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := heroCollection.EnsureIndex(index); err != nil {
		panic(err)
	}
}

func (hero *Hero) CreateHero() error {
	hero._ID = bson.NewObjectId()
	c := DB.C("heroes")
	return c.Insert(hero)
}

func GetHeroes() ([]Hero, error) {
	heroes := []Hero{}
	c := DB.C("heroes")
	return heroes, c.Find(nil).All(&heroes)
}

func GetHero(id int) (*Hero, error) {
	hero := Hero{}
	c := DB.C("heroes")
	return &hero, c.Find(bson.M{"id": id}).One(&hero)
}

func ModifyHero(hero Hero) (*Hero, error) {
	c := DB.C("heroes")

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"name": hero.Name}},
		ReturnNew: true,
	}

	_, err := c.Find(bson.M{"id": hero.ID}).Apply(change, &hero)
	return &hero, err
}

func RemoveHero(id int) error {
	c := DB.C("heroes")
	return c.Remove(bson.M{"id": id})
}
