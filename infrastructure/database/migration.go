package database

import (
	"github.com/VooDooStack/FitStackAPI/domain/exercise"
	healthLogs "github.com/VooDooStack/FitStackAPI/domain/health_logs"
	"github.com/VooDooStack/FitStackAPI/domain/program"
	"github.com/VooDooStack/FitStackAPI/domain/routine"
	"github.com/VooDooStack/FitStackAPI/domain/user"

	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&program.Program{},
		&user.User{},
		&user.UserProfile{},
		&user.UserStatistic{},
		&healthLogs.BodyFatPercentageLog{},
		&healthLogs.BodyMassIndexLog{},
		&healthLogs.WeightLog{},
		&healthLogs.HeightLog{},
		&healthLogs.StepsLog{},
		&healthLogs.ActiveEnergyBurnedLog{},
		&healthLogs.BasalEnergyBurnedLog{},
		&healthLogs.HeartRateLog{},
		&healthLogs.SleepAsleepLog{},
		&healthLogs.SleepAwakeLog{},
		&healthLogs.SleepInBedLog{},
		&user.Achievement{},
		&user.Challenge{},
		&user.Friendship{},
		&exercise.ExerciseEquipment{},
		&exercise.MuscleTarget{},
		&exercise.Exercise{},
		&exercise.WorkoutSets{},
		&exercise.Workout{},
		&routine.RoutineSchedule{},
		&routine.Routine{},
	)
	if err != nil {
		return err
	}

	return nil
}
