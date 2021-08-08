package biz

import (
	bizmodels "HMThird/internal/biz/bizModels"
	"log"
)

type UserAccountRepo interface {
	CreateUserAccount(*bizmodels.UserAccount) error
	UpdateUserAccount(*bizmodels.UserAccount) error
}

type UserAccountUsercase struct {
	repo UserAccountRepo
	log  log.Logger
}

func NewUserAccountUsecase(repo UserAccountRepo, logger log.Logger) *UserAccountUsercase {
	return &UserAccountUsercase{repo, logger}
}

func (uc *UserAccountUsercase) Create(account *bizmodels.UserAccount) error {
	return uc.repo.CreateUserAccount(account)
}

func (uc *UserAccountUsercase) Update(account *bizmodels.UserAccount) error {
	return uc.repo.UpdateUserAccount(account)
}
