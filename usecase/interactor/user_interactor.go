package interactor

import (
	"github.com/Rasukarusan/go-redmine-time-entry/domain"
	"github.com/Rasukarusan/go-redmine-time-entry/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
}

func (interactor *UserInteractor) UserLogin(u domain.User) (bool, error) {
	return interactor.UserRepository.Login(u)
}
