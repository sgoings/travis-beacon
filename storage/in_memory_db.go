package storage

import (
	"log"

	"github.com/sgoings/travis-beacon/chart"
)

// inMemoryDB supplies an implementation of Db
type inMemoryDB struct {
	charts map[string]chart.Chart
}

// NewInMemoryDB creates a new InMemoryDB
func NewInMemoryDB() DB {
	return &inMemoryDB{charts: make(map[string]chart.Chart)}
}

//Get returns a chart by name
func (db *inMemoryDB) Get(name string) chart.Chart {
	return db.charts[name]
}

//GetAll returns all charts
func (db *inMemoryDB) GetAll() map[string]chart.Chart {
	return db.charts
}

//Set updates an existing chart or creates a new one
func (db *inMemoryDB) Set(c chart.Chart) {
	db.charts[c.Name] = c
	log.Printf("[DEBUG] Updating chart '%s' (%d/100) in DB\n", c.Name, c.Status)
}
