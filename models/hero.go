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

func InitHeroCollection(db *mgo.Database) {
	heroCollection = db.C("heroes")

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
	return heroCollection.Insert(hero)
}

func (hero *Hero) IsIn(heroes []Hero) bool {
	for _, s := range heroes {
		if s.ID == hero.ID && s.Name == hero.Name {
			return true
		}
	}

	return false
}

func (hero *Hero) Equals(other Hero) bool {
	return hero.ID == other.ID && hero.Name == other.Name
}

func GetHeroes() ([]Hero, error) {
	heroes := []Hero{}
	return heroes, heroCollection.Find(nil).All(&heroes)
}

func GetHero(id int) (*Hero, error) {
	hero := Hero{}
	return &hero, heroCollection.Find(bson.M{"id": id}).One(&hero)
}

func ModifyHero(hero Hero) (*Hero, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"name": hero.Name}},
		ReturnNew: true,
	}

	_, err := heroCollection.Find(bson.M{"id": hero.ID}).Apply(change, &hero)
	return &hero, err
}

func RemoveHero(id int) error {
	return heroCollection.Remove(bson.M{"id": id})
}
