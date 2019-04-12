package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Hero struct {
	ID   int    `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

var heroCollection *mgo.Collection

func InitHeroCollection(db *mgo.Database) {
	heroCollection = db.C("heroes")
}

func (hero *Hero) CreateHero() error {
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
	return &hero, heroCollection.Find(bson.M{"_id": id}).One(&hero)
}

func ModifyHero(hero Hero) (*Hero, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"name": hero.Name}},
		ReturnNew: true,
	}

	_, err := heroCollection.Find(bson.M{"_id": hero.ID}).Apply(change, &hero)
	return &hero, err
}

func RemoveHero(id int) error {
	return heroCollection.Remove(bson.M{"_id": id})
}
