package controller

import (
	"log"

	"github.com/Rasukarusan/go-redmine-time-entry/domain"
	"github.com/Rasukarusan/go-redmine-time-entry/infrastructure"
	"github.com/Rasukarusan/go-redmine-time-entry/interfaces/browser"
	"github.com/Rasukarusan/go-redmine-time-entry/usecase/interactor"
	"github.com/sclevine/agouti"
)

type UserController struct {
	Interactor interactor.UserInteractor
}

func NewUserController(driver *agouti.WebDriver) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: &browser.UserRepository{
				Browser: infrastructure.NewBrowser(driver),
			},
		},
	}
}

func (controller *UserController) SignIn() {
	u := domain.User{
		Username: "tanaka.naoto",
		Password: "passooowoww",
	}
	isLogin, err := controller.Interactor.UserLogin(u)
	if err != nil {
		log.Fatal(err)
	}
	if isLogin {
		log.Println("Login Success!")
	}
}
