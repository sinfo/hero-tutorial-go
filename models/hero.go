package models

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Hero represents a hero to be stored in the database and parsed from http requests
type Hero struct {
	ID   int    `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

var heroCollection *mgo.Collection

// InitHeroCollection grabs the hero collection from the database
func InitHeroCollection(db *mgo.Database) {
	heroCollection = db.C("heroes")
}

// HeroFromBody parses the body of an http request to a hero and validates it
func HeroFromBody(body io.Reader) (*Hero, error) {
	var hero Hero

	if err := json.NewDecoder(body).Decode(&hero); err != nil {
		return nil, err
	}

	if len(hero.Name) == 0 {
		return nil, errors.New("invalid name on hero")
	}

	return &hero, nil
}

// CreateHero saves the hero to the database
func (hero *Hero) CreateHero() error {
	return heroCollection.Insert(hero)
}

// IsIn checks if the hero can be found on the slice of heroes given as input
func (hero *Hero) IsIn(heroes []Hero) bool {
	for _, s := range heroes {
		if s.ID == hero.ID && s.Name == hero.Name {
			return true
		}
	}

	return false
}

// Equals compares 2 heroes and checks if they are equal in value
func (hero *Hero) Equals(other Hero) bool {
	return hero.ID == other.ID && hero.Name == other.Name
}

// GetHeroes gets all heroes from the database
func GetHeroes() ([]Hero, error) {
	heroes := []Hero{}
	return heroes, heroCollection.Find(nil).All(&heroes)
}

// GetHero gets a hero from the database, using its ID
func GetHero(id int) (*Hero, error) {
	hero := Hero{}
	return &hero, heroCollection.Find(bson.M{"_id": id}).One(&hero)
}

// ModifyHero uses the hero given in the arguments to
// find a hero with equal ID on the database, and change
// its name to equal the name of the hero given in the arguments
func ModifyHero(hero Hero) (*Hero, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"name": hero.Name}},
		ReturnNew: true,
	}

	_, err := heroCollection.Find(bson.M{"_id": hero.ID}).Apply(change, &hero)
	return &hero, err
}

// RemoveHero deletes a hero from the database
func RemoveHero(id int) error {
	return heroCollection.Remove(bson.M{"_id": id})
}
