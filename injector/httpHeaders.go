package injector 

import (
	"fmt"
	"time"
	"net/http"
)

func formatHeaderDate(t time.Time) string {
	return t.UTC().Format(http.TimeFormat)
}

func generateHttpHeaders(
	contentSize int, 
	contentType string,
	currDate time.Time, 
	expireDate time.Time,
) []string {
	currDateStr := formatHeaderDate(currDate)
	headers := []string {
		fmt.Sprintf("Content-Length: %v", contentSize),
		fmt.Sprintf("Content-Type: %v; charset=utf-8", contentType),
		"Expires: " + formatHeaderDate(expireDate),
		"Last-Modified: " + currDateStr,
		"Date: " + currDateStr,
	}
	return headers
}