package platform

import (
	"path/filepath"
)

func pathsForPatterns(patterns []string) []string {
	results := make(map[string]bool, 0)
	
	for _, pattern := range patterns {
		if paths, _ := filepath.Glob(pattern); len(paths) > 0 {

			for _, path := range paths {
				results[path] = true 
			}
		}
	}
	uniqueResults := make([]string, 0)
	for path := range results {
		uniqueResults = append(uniqueResults, path)
 	}
	return uniqueResults
}

func CacheDirs() []string {
	patterns := cacheDirPatterns()
	return pathsForPatterns(patterns)
}