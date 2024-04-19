package main 

import (
	"fmt"

	"chrome-poc/injector"
)

func main() {
	fmt.Print("Chrome POC demo")

	chromeInjector := injector.ChromeInjector{}
	chromeInjector.Inject()
}