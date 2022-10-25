package exercise

type Exercise struct {
	Name              string
	Description       string
	Image             string
	MetValue          string
	ExerciseType      []*ExerciseType
	ExerciseEquipment []*ExerciseEquipment
	MuscleTarget      []*MuscleTarget
}
