package user

import (
	"time"

	healthLogs "github.com/VooDooStack/FitStackAPI/domain/health_logs"
	"gorm.io/gorm"
)

type UserStatistic struct {
	ID                     string                             `json:"id" gorm:"primaryKey;"`
	WeightLogs             []healthLogs.WeightLog             `json:"weight_logs" gorm:"foreignKey:UserStatisticID"`
	BodyMassIndexLogs      []healthLogs.BodyMassIndexLog      `json:"body_mass_index_logs" gorm:"foreignKey:UserStatisticID"`
	BodyFatPercentageLogs  []healthLogs.BodyFatPercentageLog  `json:"body_fat_percentage_logs" gorm:"foreignKey:UserStatisticID"`
	StepsLogs              []healthLogs.StepsLog              `json:"steps_logs" gorm:"foreignKey:UserStatisticID"`
	HeartRateLogs          []healthLogs.HeartRateLog          `json:"heart_rate_logs" gorm:"foreignKey:UserStatisticID"`
	SleepAsleepLogs        []healthLogs.SleepAsleepLog        `json:"sleep_asleep_logs" gorm:"foreignKey:UserStatisticID"`
	SleepAwakeLogs         []healthLogs.SleepAwakeLog         `json:"sleep_awake_logs" gorm:"foreignKey:UserStatisticID"`
	SleepInBedLogs         []healthLogs.SleepInBedLog         `json:"sleep_in_bed_logs" gorm:"foreignKey:UserStatisticID"`
	ActiveEnergyBurnedLogs []healthLogs.ActiveEnergyBurnedLog `json:"active_energy_burned_logs" gorm:"foreignKey:UserStatisticID"`
	BasalEnergyBurnedLog   []healthLogs.BasalEnergyBurnedLog  `json:"basal_energy_burned_logs" gorm:"foreignKey:UserStatisticID"`
	CreatedAt              time.Time                          `json:"created_at"`
	UpdatedAt              time.Time                          `json:"updated_at"`
	DeletedAt              gorm.DeletedAt                     `json:"deleted_at" gorm:"index"`
}
