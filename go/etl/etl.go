package etl

import (
	"strings"
)

// Transforms the old map provided into a new map
func Transform(oldMap map[int][]string) map[string]int {

	newMap := map[string]int{}

	for pts, coll := range oldMap {
		// old key is new value
		// each old value are new keys
		for _, let := range coll {
			newMap[strings.ToLower(let)] = pts
		}

	}

	return newMap
}
