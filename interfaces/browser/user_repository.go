package browser

import (
	"log"

	"github.com/Rasukarusan/go-redmine-time-entry/domain"
)

type UserRepository struct {
	Browser Browser
}

func (repo *UserRepository) Login(u domain.User) (bool, error) {
	log.Println(u.Username)
	log.Println(u.Password)
	return true, nil
}
