package usecase

import (
	"github.com/Uchel/auth-jwt2/repository"
)

type ExampleAdminLoginUsecase interface {
	FindByEmailExampleAdmin(email string) (string, string)
}

type exampleAdminLoginUsecase struct {
	exampleAdminRepo repository.ExampleAdminLoginRepo
}

//==========================================================================================

func (u *exampleAdminLoginUsecase) FindByEmailExampleAdmin(email string) (string, string) {

	return u.exampleAdminRepo.GetByEmailExampleAdmin(email)
}

func NewExampleAdminLoginUsecase(exampleAdminLoginRepo repository.ExampleAdminLoginRepo) ExampleAdminLoginUsecase {
	return &exampleAdminLoginUsecase{
		exampleAdminRepo: exampleAdminLoginRepo,
	}
}
