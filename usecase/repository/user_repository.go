package repository

import "github.com/Rasukarusan/go-redmine-time-entry/domain"

type UserRepository interface {
	Login(domain.User) (bool, error)
}
