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
	for _, pageUrl := range self.PageUrls {
		// TODO: remove test file
		filename, fileData := self.generatePayload(pageUrl)
		path := fmt.Sprintf("./test-files/%s", filename)
		ioutil.WriteFile(path, fileData, 0644)
	}
}

func (self *ChromeInjector) generatePayload(urlStr string) (string, []byte) {
	pageUrl, _ := url.Parse(urlStr)
	cacheKey := generateCacheKey(pageUrl)
	eHash := entryHash(cacheKey)
	filename := self.getFilename(eHash, 0)

	fileData := make([]byte, 0)
	fileData = append(fileData, cacheKey...)
	return filename, fileData
}

func (self *ChromeInjector) getFilename(
    entryHash uint64,
    fileIndex int,
) string {
  	filename := fmt.Sprintf(
		"%x_%1d", 
		entryHash, 
		fileIndex,
  	)
	return filename
}