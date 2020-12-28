package model

import (
	"strings"
)

// csvToMap splits a comma-separated list into a map
func csvToMap(s string) map[string]int {

	l := make(map[string]int)

	for i, v := range strings.Split(s, ",") {
		l[v] = i
	}
	return l
}
