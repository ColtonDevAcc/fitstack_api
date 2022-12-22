package muscle

type RecoveryMuscle struct {
	RecoveryUserID string `json:"-" gorm:"primaryKey; unique"`
	MuscleID       uint   `json:"-" gorm:"primaryKey; unique"`
	RecoveryValue  int    `json:"recovery_value" gorm:"default:0; not null;"`
}
