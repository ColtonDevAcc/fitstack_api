package exercise

type Exercise struct {
	Name              string              `json:"name" db:"name"`
	Description       string              `json:"description" db:"description"`
	Image             *string             `json:"image" db:"image"`
	MetValue          float32             `json:"met_value" db:"met_value"`
	ExerciseType      []ExerciseType      `db:""`
	ExerciseEquipment []ExerciseEquipment `db:""`
	MuscleTarget      []MuscleTarget      `db:""`
}
