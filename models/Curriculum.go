package models

type Curriculum struct {
	ID              string `gorm:"type:VARCHAR(2)"`
	Program_Code    string `gorm:"type:VARCHAR(10)"`
	Revision_Number uint   `gorm:"type: INT UNSIGNED"`
	Effectivity_Sem uint   `gorm:"type: INT UNSIGNED"`
	Effectivity_SY  uint   `gorm:"type: INT UNSIGNED"`
	CMO_Ref         string `gorm:"type:VARCHAR(255)"`
	IsActive        bool
}
