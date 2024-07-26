package models

type Curriculum struct {
	ID              string
	Revision_Number uint
	Effectivity_Sem uint
	Effectivity_SY  uint
	CMO_Ref         string
	IsActive        bool
}
