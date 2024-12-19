package core

import "strings"

func SanitizeSuffixPath(path string) string {
	if len(path) == 1 {
		return path
	}
	return strings.TrimSuffix(path, "/")
}
