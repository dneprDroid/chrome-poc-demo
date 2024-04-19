package injector 

import (
	"fmt"
	"net/url"
	"io/ioutil"
)

type ChromeInjector struct {
	PageUrls []string
}

func (self *ChromeInjector) Inject() {
	for i, pageUrl := range self.PageUrls {
		// TODO: remove test file
		path := fmt.Sprintf("./test-file-%i", i)
		fileData := self.generatePayload(pageUrl)
		ioutil.WriteFile(path, fileData, 0644)
	}
}

func (self *ChromeInjector) generatePayload(urlStr string) []byte {
	pageUrl, _ := url.Parse(urlStr)
	cacheKey := generateCacheKey(pageUrl)

	fileData := make([]byte, 0)
	fileData = append(fileData, cacheKey...)
	return fileData
}