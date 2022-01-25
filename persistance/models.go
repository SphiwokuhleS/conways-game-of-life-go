package persistance

import "gorm.io/gorm"

type World struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Grid  string
	Epoch int
}
