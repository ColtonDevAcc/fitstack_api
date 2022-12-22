package muscle

import (
	"time"

	"gorm.io/gorm"
)

type MuscleGroup string

type Muscle struct {
	ID        uint           `json:"id" gorm:"primaryKey; autoIncrement;"`
	Action    string         `json:"action"`
	Group     MuscleGroup    `json:"group" gorm:"type:muscle_group;not null;uniqueIndex:idx_muscle_group_name;"`
	Child     ChildMuscle    `json:"child"  gorm:"type:child_muscle; not null;uniqueIndex:idx_muscle_group_name;"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Recovery  RecoveryMuscle `json:"recovery" gorm:"-"`
}

const (
	Arms           MuscleGroup = "Arms"
	Legs           MuscleGroup = "Legs"
	Abdominals     MuscleGroup = "Abdominals"
	Glutes         MuscleGroup = "Glutes"
	Calves         MuscleGroup = "Calves"
	Forearms       MuscleGroup = "Forearms"
	Traps          MuscleGroup = "Traps"
	Neck           MuscleGroup = "Neck"
	UpperBack      MuscleGroup = "UpperBack"
	LowerBack      MuscleGroup = "LowerBack"
	HipFlexors     MuscleGroup = "HipFlexors"
	Adductors      MuscleGroup = "Adductors"
	Abductors      MuscleGroup = "Abductors"
	PlantarFlexors MuscleGroup = "PlantarFlexors"
	Dorsiflexors   MuscleGroup = "Dorsiflexors"
	Invertors      MuscleGroup = "Invertors"
)

type ChildMuscle string

const (
	Biceps                  ChildMuscle = "Biceps"
	Triceps                 ChildMuscle = "Triceps"
	Quadriceps              ChildMuscle = "Quadriceps"
	Hamstrings              ChildMuscle = "Hamstrings"
	RectusAbdominis         ChildMuscle = "RectusAbdominis"
	Obliques                ChildMuscle = "Obliques"
	GluteusMaximus          ChildMuscle = "GluteusMaximus"
	GluteusMedius           ChildMuscle = "GluteusMedius"
	GluteusMinimus          ChildMuscle = "GluteusMinimus"
	Gastrocnemius           ChildMuscle = "Gastrocnemius"
	Soleus                  ChildMuscle = "Soleus"
	Flexors                 ChildMuscle = "Flexors"
	Extensors               ChildMuscle = "Extensors"
	Sternocleidomastoid     ChildMuscle = "Sternocleidomastoid"
	Splenius                ChildMuscle = "Splenius"
	Rhomboids               ChildMuscle = "Rhomboids"
	LatissimusDorsi         ChildMuscle = "LatissimusDorsi"
	ErectorSpinae           ChildMuscle = "ErectorSpinae"
	Iliopsoas               ChildMuscle = "Iliopsoas"
	RectusFemoris           ChildMuscle = "RectusFemoris"
	Gracilis                ChildMuscle = "Gracilis"
	AdductorBrevis          ChildMuscle = "AdductorBrevis"
	AdductorLongus          ChildMuscle = "AdductorLongus"
	GluteusMediusMinimus    ChildMuscle = "GluteusMediusMinimus"
	TensorFasciaeLatae      ChildMuscle = "TensorFasciaeLatae"
	TibialisAnterior        ChildMuscle = "TibialisAnterior"
	TibialisPosterior       ChildMuscle = "TibialisPosterior"
	ExtensorDigitorumLongus ChildMuscle = "ExtensorDigitorumLongus"
	FlexorDigitorumLongus   ChildMuscle = "FlexorDigitorumLongus"
	PectoralisMajor         ChildMuscle = "PectoralisMajor"
	PectoralisMinor         ChildMuscle = "PectoralisMinor"
)

// Chest muscles: These muscles include the pectoralis major and minor, which are responsible for moving the arms and shoulders.

// Back muscles: These muscles include the latissimus dorsi, trapezius, and erector spinae, which are responsible for pulling, lifting, and stabilizing the body.

// Shoulder muscles: These muscles include the deltoids, rotator cuff, and trapezius, which are responsible for lifting, rotating, and stabilizing the arms.

// Arm muscles: These muscles include the biceps, triceps, and forearms, which are responsible for flexing, extending, and rotating the arms.

// Leg muscles: These muscles include the quadriceps, hamstrings, and calves, which are responsible for extending, flexing, and stabilizing the legs.

// Core muscles: These muscles include the abdominal muscles, obliques, and lower back muscles, which are responsible for stabilizing and supporting the spine and pelvis.

// Within each of these muscle groups, there are also smaller muscles that work together to perform specific movements. For example, the chest muscles include the pectoralis major, which is responsible for moving the arms, and the pectoralis minor, which is responsible for stabilizing the shoulders. The back muscles include the latissimus dorsi, which is responsible for pulling, and the trapezius, which is responsible for lifting and stabilizing the shoulders.
