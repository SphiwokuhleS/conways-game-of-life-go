package persistance

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func CreateWorld(world World) World {
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}
	db.Create(&world)
	return world
}

func GetWorldByName(worldName string) World {
	var world World
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}

	db.Where(map[string]interface{}{"Name": worldName}).Find(&world)
	return world
}

func GetAllWorlds() []World {
	var worlds []World
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}

	db.Model(&World{}).Find(&worlds)
	return worlds
}

func GetWorldEpoch(worldName string) int {
	var world World
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}

	db.Where(map[string]interface{}{"Name": worldName}).Find(&world)
	return world.Epoch
}

func UpdateEpochWorldEpoch(worldName string) {
	var world World
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}

	db.Where(map[string]interface{}{"Name": worldName}).Find(&world)
	newEpoch := world.Epoch + 1
	db.Model(&world).Update("Epoch", newEpoch)
}

func UpdateGridWorldGrid(worldName string, newGrid string) {
	var world World
	db, err := connect()

	if err != nil {
		log.Println("Could not connect to database", err)
	}

	db.Where(map[string]interface{}{"Name": worldName}).Find(&world)
	db.Model(&world).Update("Grid", newGrid)

}

func FindWorldById() {
	//Delete World
}

func DeleteWorld(world World) bool {
	db, err := connect()
	if err != nil {
		log.Println("Could not connect to database", err)
		return false
	}

	db.Delete(&world)
	return true
}

func connect() (*gorm.DB, error) {
	//connect to database
	db, err := gorm.Open(sqlite.Open("conways.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, err
}
