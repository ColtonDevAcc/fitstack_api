package exercise

type Exercise struct {
	Name              string              `db:"name"`
	Description       string              `db:"description"`
	Image             string              `db:"image"`
	MetValue          string              `db:"met_value"`
	ExerciseType      []ExerciseType      `db:""`
	ExerciseEquipment []ExerciseEquipment `db:""`
	MuscleTarget      []MuscleTarget      `db:""`
}
