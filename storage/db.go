package storage

import (
	"github.com/sgoings/travis-beacon/chart"
)

// DB is the interface to get/store data
type DB interface {
	Get(name string) chart.Chart
	GetAll() map[string]chart.Chart
	Set(chart.Chart)
}
