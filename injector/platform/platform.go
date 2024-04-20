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

	pathResults := make([]string, 0)
	for path := range results {
        pathResults = append(pathResults, path)
    }
	return pathResults
}

func CacheDirs() []string {
	patterns := cacheDirPatterns()
	return pathsForPatterns(patterns)
}