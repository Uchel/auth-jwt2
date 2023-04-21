package usecase

import (
	"github.com/Uchel/auth-jwt2/repository"
)

type IcTeamLoginUsecase interface {
	FindByEmailIc(email string) (string, string)
}

type icTeamLoginUsecase struct {
	icTeamLoginRepo repository.IcTeamLoginRepo
}

//==========================================================================================

func (u icTeamLoginUsecase) FindByEmailIc(email string) (string, string) {

	return u.icTeamLoginRepo.GetByEmailIc(email)
}

func NewUserUsecase(icTeamLoginRepo repository.IcTeamLoginRepo) IcTeamLoginUsecase {
	return &icTeamLoginUsecase{
		icTeamLoginRepo: icTeamLoginRepo,
	}
}
