package injector

import (
	"strings"
	"unicode"
	"fmt"
	"net/url"
)

func generateCacheKey(
	pageUrl *url.URL,
) string {
	credKey := "1"
	postKey := "0"
	schemeWithDomain := pageUrl.Scheme + "://" + urlDomain(pageUrl)
	key := fmt.Sprintf(
		"%s/%s/_dk_%s %s %s",
		credKey, postKey, 
		schemeWithDomain, schemeWithDomain,
		pageUrl.String(),
	)
	return key 
}

func urlDomain(pageUrl *url.URL) string {
	hostname := pageUrl.Hostname()
	parts := strings.Split(hostname, ".")
	if len(parts) <= 2 {
		return hostname
	}
	lastPart := parts[len(parts)-1]
	if len(lastPart) == 1 && 
	   !unicode.IsLetter([]rune(lastPart)[0]) {
		return hostname
	}
	domain := parts[len(parts)-2] + "." + lastPart
	return domain
}