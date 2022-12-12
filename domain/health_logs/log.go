package healthLogs

import (
	"time"

	"gorm.io/gorm"
)

// create interface for logs
type Log struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Value           float32        `json:"value"`
}

// create enum for health data types
type HealthDataType string

const (
	ACTIVE_ENERGY_BURNED        HealthDataType = "ACTIVE_ENERGY_BURNED"
	AUDIOGRAM                   HealthDataType = "AUDIOGRAM"
	BASAL_ENERGY_BURNED         HealthDataType = "BASAL_ENERGY_BURNED"
	BLOOD_GLUCOSE               HealthDataType = "BLOOD_GLUCOSE"
	BLOOD_OXYGEN                HealthDataType = "BLOOD_OXYGEN"
	BLOOD_PRESSURE_DIASTOLIC    HealthDataType = "BLOOD_PRESSURE_DIASTOLIC"
	BLOOD_PRESSURE_SYSTOLIC     HealthDataType = "BLOOD_PRESSURE_SYSTOLIC"
	BODY_FAT_PERCENTAGE         HealthDataType = "BODY_FAT_PERCENTAGE"
	BODY_MASS_INDEX             HealthDataType = "BODY_MASS_INDEX"
	BODY_TEMPERATURE            HealthDataType = "BODY_TEMPERATURE"
	DIETARY_CARBS_CONSUMED      HealthDataType = "DIETARY_CARBS_CONSUMED"
	DIETARY_ENERGY_CONSUMED     HealthDataType = "DIETARY_ENERGY_CONSUMED"
	DIETARY_FATS_CONSUMED       HealthDataType = "DIETARY_FATS_CONSUMED"
	DIETARY_PROTEIN_CONSUMED    HealthDataType = "DIETARY_PROTEIN_CONSUMED"
	FORCED_EXPIRATORY_VOLUME    HealthDataType = "FORCED_EXPIRATORY_VOLUME"
	HEART_RATE                  HealthDataType = "HEART_RATE"
	HEART_RATE_VARIABILITY_SDNN HealthDataType = "HEART_RATE_VARIABILITY_SDNN"
	HEIGHT                      HealthDataType = "HEIGHT"
	RESTING_HEART_RATE          HealthDataType = "RESTING_HEART_RATE"
	STEPS                       HealthDataType = "STEPS"
	WAIST_CIRCUMFERENCE         HealthDataType = "WAIST_CIRCUMFERENCE"
	WALKING_HEART_RATE          HealthDataType = "WALKING_HEART_RATE"
	WEIGHT                      HealthDataType = "WEIGHT"
	DISTANCE_WALKING_RUNNING    HealthDataType = "DISTANCE_WALKING_RUNNING"
	FLIGHTS_CLIMBED             HealthDataType = "FLIGHTS_CLIMBED"
	MOVE_MINUTES                HealthDataType = "MOVE_MINUTES"
	DISTANCE_DELTA              HealthDataType = "DISTANCE_DELTA"
	MINDFULNESS                 HealthDataType = "MINDFULNESS"
	WATER                       HealthDataType = "WATER"
	SLEEP_IN_BED                HealthDataType = "SLEEP_IN_BED"
	SLEEP_ASLEEP                HealthDataType = "SLEEP_ASLEEP"
	SLEEP_AWAKE                 HealthDataType = "SLEEP_AWAKE"
	EXERCISE_TIME               HealthDataType = "EXERCISE_TIME"
	WORKOUT                     HealthDataType = "WORKOUT"
	HEADACHE_NOT_PRESENT        HealthDataType = "HEADACHE_NOT_PRESENT"
	HEADACHE_MILD               HealthDataType = "HEADACHE_MILD"
	HEADACHE_MODERATE           HealthDataType = "HEADACHE_MODERATE"
	HEADACHE_SEVERE             HealthDataType = "HEADACHE_SEVERE"
	HEADACHE_UNSPECIFIED        HealthDataType = "HEADACHE_UNSPECIFIED"

	// Heart Rate events (specific to Apple Watch)
	HIGH_HEART_RATE_EVENT      HealthDataType = "HIGH_HEART_RATE_EVENT"
	LOW_HEART_RATE_EVENT       HealthDataType = "LOW_HEART_RATE_EVENT"
	IRREGULAR_HEART_RATE_EVENT HealthDataType = "IRREGULAR_HEART_RATE_EVENT"
	ELECTRODERMAL_ACTIVITY     HealthDataType = "ELECTRODERMAL_ACTIVITY"
)
