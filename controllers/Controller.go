package controllers

type Controller struct {
	Auth       *AuthController
	User       *UserController
	Curriculum *CurriculumController
}
