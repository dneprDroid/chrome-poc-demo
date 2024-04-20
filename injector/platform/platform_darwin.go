package platform

func cacheDirPatterns() []string {
	return []string {
		"/Users/*/Library/Caches/Google/*/*/Cache/Cache_Data",
		"/Users/*/Library/Caches/*/*/*/Cache/Cache_Data",
		"/Users/*/Library/Caches/*/*/Cache/Cache_Data",
	}
}