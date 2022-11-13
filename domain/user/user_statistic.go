package user

import (
	"time"

	"gorm.io/gorm"
)

type UserStatistic struct {
	ID                   string                `json:"id" gorm:"primaryKey;"`
	HeightLogs           []HeightLog           `json:"height_log" gorm:"foreignKey:UserStatisticID"`
	WeightLogs           []WeightLog           `json:"weight_log" gorm:"foreignKey:UserStatisticID"`
	BMILogs              []BMILog              `json:"bmi_log" gorm:"foreignKey:UserStatisticID"`
	BodyFatLogs          []BodyFatLog          `json:"body_fat_log" gorm:"foreignKey:UserStatisticID"`
	StepLogs             []StepsLog            `json:"step_log" gorm:"foreignKey:UserStatisticID"`
	DistanceLogs         []DistanceLog         `json:"distance_logs" gorm:"foreignKey:UserStatisticID"`
	ActiveMinutesLogs    []ActiveMinutesLog    `json:"active_minutes_logs" gorm:"foreignKey:UserStatisticID"`
	HeartRateLogs        []HeartRateLog        `json:"heart_rate_logs" gorm:"foreignKey:UserStatisticID"`
	SleepLogs            []SleepLog            `json:"sleep_logs" gorm:"foreignKey:UserStatisticID"`
	BloodPressureLogs    []BloodPressureLog    `json:"blood_pressure_logs" gorm:"foreignKey:UserStatisticID"`
	BodyTemperatureLogs  []TemperatureLog      `json:"body_temperature_logs" gorm:"foreignKey:UserStatisticID"`
	RespiratoryRateLogs  []RespiratoryLog      `json:"respiratory_rate_logs" gorm:"foreignKey:UserStatisticID"`
	OxygenSaturationLogs []OxygenSaturationLog `json:"oxygen_saturation_logs" gorm:"foreignKey:UserStatisticID"`
	ActiveEnergyLogs     []ActiveEnergyLog     `json:"active_energy_logs" gorm:"foreignKey:UserStatisticID"`
	BasalEnergyLog       []BasalEnergyLog      `json:"basal_energy_logs" gorm:"foreignKey:UserStatisticID"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	DeletedAt            gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
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

type StepsLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Steps           float32        `json:"steps"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type DistanceLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Distance        float32        `json:"distance"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ActiveMinutesLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	ActiveMinutes   float32        `json:"active_minutes"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type SleepLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Sleep           float32        `json:"sleep"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type HeartRateLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	HeartRate       float32        `json:"heart_rate"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BloodPressureLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Systolic        float32        `json:"systolic"`
	Diastolic       float32        `json:"diastolic"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BloodGlucoseLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	BloodGlucose    float32        `json:"blood_glucose"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type OxygenSaturationLog struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID  string         `json:"user_statistic_id" gorm:"not null"`
	OxygenSaturation float32        `json:"oxygen_saturation"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TemperatureLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Temperature     float32        `json:"temperature"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type LeanBodyMassLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	LeanBodyMass    float32        `json:"lean_body_mass"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ActiveEnergyLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	ActiveEnergy    float32        `json:"active_energy"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type BasalEnergyLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	BasalEnergy     float32        `json:"basal_energy"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RunningSpeedLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	RunningSpeed    float32        `json:"running_speed"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RunningCadenceLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	RunningCadence  float32        `json:"running_cadence"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type WalkingSpeedLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	WalkingSpeed    float32        `json:"walking_speed"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type WalkingCadenceLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	WalkingCadence  float32        `json:"walking_cadence"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CyclingSpeedLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	CyclingSpeed    float32        `json:"cycling_speed"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CyclingCadenceLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	CyclingCadence  float32        `json:"cycling_cadence"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RespiratoryLog struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserStatisticID string         `json:"user_statistic_id" gorm:"not null"`
	Respiratory     float32        `json:"respiratory"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
