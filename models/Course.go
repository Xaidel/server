package models

type Course struct{
  ID uint `gorm:"type: INT UNSIGNED"`
  CurriculumID string `gorm:"type: VARCHAR(2)"`
  Curriculum Curriculum
  Course_No string `gorm:"type: VARCHAR(255)"`
  Course_Desc string `gorm:"type: VARCHAR(255)"`
  Lecture_Unit uint `gorm:"type: INT UNSIGNED"`
  Semester uint `gorm:"type: INT UNSIGNED"`
  Year_Level uint `gorm:"type: INT UNSIGNED"`
}
