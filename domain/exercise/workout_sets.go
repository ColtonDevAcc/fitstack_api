package exercise

import "gorm.io/gorm"

type WorkoutSets struct {
	gorm.Model
	Exercise *Exercise `json:"exercise" gorm:"foreignKey:id;references:id"`
}
