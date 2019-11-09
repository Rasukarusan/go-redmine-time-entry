package main

import (
	"github.com/Rasukarusan/go-redmine-time-entry/infrastructure"
	"github.com/Rasukarusan/go-redmine-time-entry/interfaces/controller"
)

func main() {
	driver := infrastructure.Driver
	driver.Start()
	defer driver.Stop()

	userController := controller.NewUserController(driver)
	userController.SignIn()

	// c := config.Load()

	//
	// // ログイン
	// page, _ := driver.NewPage(agouti.Browser("chrome"))
	// page.Navigate(c.Url + "/login")
	// page.FindByID("username").SendKeys(c.UserName)
	// page.FindByID("password").SendKeys(c.Password)
	// page.FindByName("login").Submit()
	//
	// // 登録内容を取得
	// timeEntry := domain.TimeEntry{
	// 	SpentOn:    "2019-11-01",
	// 	Hour:       "hoge",
	// 	Comment:    "hoge",
	// 	ActivityID: "調査作業",
	// }
	//
	// // 時間記録
	// page.Navigate(c.Url + "/issues/59865/time_entries/new")
	// page.FindByID("")
	// page.FindByID("time_entry_spent_on").Clear()
	// page.FindByID("time_entry_spent_on").SendKeys(timeEntry.SpentOn)
	// page.FindByID("time_entry_hours").SendKeys(timeEntry.Hour)
	// page.FindByID("time_entry_comments").SendKeys(timeEntry.Comment)
	// page.FindByID("time_entry_activity_id").Select(timeEntry.ActivityID)
	// page.FindByName("commit").Click()

}
