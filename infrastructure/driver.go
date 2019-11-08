package infrastructure

import (
	"github.com/sclevine/agouti"
)

var Driver *agouti.WebDriver

func init() {
	options := agouti.ChromeOptions("args", []string{
		"--headless",
		"--window-size=1280,800",
	})
	driver := agouti.ChromeDriver(options)
	Driver = driver
}
