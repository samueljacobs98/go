package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// option 1 sync.Once
// option 2 init()
// consider thread safety and laziness

var once sync.Once
var instance *singletonDatabase

// func GetSingletonDatabase() *singletonDatabase {
// 	once.Do(func() {
// 		// ...
// 		caps, e := readData(".\\capitals.txt")
// 		db := singletonDatabase{}
// 		if e == nil {
// 			db.capitals = caps
// 		}
// 		instance = &db
// 	})
// 	return instance
// }

// func GetTotalPopulation(cities []string) int {
// 	result := 0
// 	for _, city := range cities {
// 		result += GetSingletonDatabase().GetPopulation(city)
// 		// DIP
// 	}
// 	return result
// }

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
		// DIP
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	// cities := []string{"Seoul", "Mexico City"}
	// // tp := GetTotalPopulation(cities)
	// tp := GetTotalPopulationEx(GetSingletonDatabase(), cities)
	// ok := tp == (17500000 + 17400000)
	// fmt.Println(ok)
	// db := GetSingletonDatabase()
	// pop := db.GetPopulation("London")
	// fmt.Println("Population of London = ", pop)

	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
