package injector

import (
	"bytes"
	"compress/gzip"
)

func gzipData(data []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func (self *ChromeInjector) GzipContent() {
	self.Content = gzipData(self.Content)
	self.ExtraHttpHeaders = append(
		self.ExtraHttpHeaders,
		"Content-Encoding: gzip",
	)
}