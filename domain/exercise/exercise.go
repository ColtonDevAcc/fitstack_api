package exercise

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name              string
	Description       string
	Image             string
	MetValue          string
	ExerciseType      []*ExerciseType      `gorm:"foreignKey:id"`
	ExerciseEquipment []*ExerciseEquipment `gorm:"foreignKey:id"`
	MuscleTarget      []*MuscleTarget      `gorm:"foreignKey:id"`
}
