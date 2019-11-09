package infrastructure

import (
	"github.com/sclevine/agouti"
)

type Browser struct {
	Page *agouti.Page
}

func NewBrowser(driver *agouti.WebDriver) *Browser {
	page, _ := driver.NewPage(agouti.Browser("chrome"))
	browser := new(Browser)
	browser.Page = page
	return browser
}
