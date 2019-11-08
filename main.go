package main

import (
	"github.com/Rasukarusan/go-redmine-time-entry/infrastructure"
	"github.com/sclevine/agouti"
)

func main() {
	driver := infrastructure.Driver
	driver.Start()
	defer driver.Stop()

	page, _ := driver.NewPage(agouti.Browser("chrome"))
	page.Navigate("http://qiita.com/")
}
