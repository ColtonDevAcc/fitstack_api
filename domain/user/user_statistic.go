package user

import (
	"time"

	"gorm.io/gorm"
)

type UserStatistic struct {
	ID         string         `json:"id" gorm:"primaryKey;"`
	HeightLog  []HeightLog    `json:"height_log" gorm:"foreignKey:UserStatisticID;"`
	WeightLog  []WeightLog    `json:"weight_log" gorm:"foreignKey:UserStatisticID;"`
	BMILog     []BMILog       `json:"bmi_log" gorm:"foreignKey:UserStatisticID;"`
	BodyFatLog []BodyFatLog   `json:"body_fat_log" gorm:"foreignKey:UserStatisticID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type HeightLog struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserStatisticID string         `json:"user_statistic_id"`
	Height          float32        `json:"height"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type WeightLog struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserStatisticID string         `json:"user_statistic_id"`
	Weight          float32        `json:"weight"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BMILog struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserStatisticID string         `json:"user_statistic_id"`
	BMI             float32        `json:"bmi"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BodyFatLog struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserStatisticID string         `json:"user_statistic_id"`
	BodyFat         float32        `json:"body_fat"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
