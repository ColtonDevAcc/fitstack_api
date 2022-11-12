package user

import (
	"time"

	"gorm.io/gorm"
)

type UserStatistic struct {
	ID          string         `json:"id" gorm:"primaryKey;"`
	HeightLogs  []HeightLog    `json:"height_log" gorm:"foreignKey:UserStatisticID"`
	WeightLogs  []WeightLog    `json:"weight_log" gorm:"foreignKey:UserStatisticID"`
	BMILogs     []BMILog       `json:"bmi_log" gorm:"foreignKey:UserStatisticID"`
	BodyFatLogs []BodyFatLog   `json:"body_fat_log" gorm:"foreignKey:UserStatisticID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type HeightLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Height          float32        `json:"height"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type WeightLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Weight          float32        `json:"weight"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BMILog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	BMI             float32        `json:"bmi"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BodyFatLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	BodyFat         float32        `json:"body_fat"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
