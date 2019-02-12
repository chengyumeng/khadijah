package hash

import (
	"strings"
	"sync"
)

type hashmap struct {
	mapping sync.Map
}

// Hashmap : Global hash map
var Hashmap hashmap

func (h *hashmap) setIfNotExist(arr []string) bool {
	_, ok := h.mapping.LoadOrStore(strings.Join(arr, ""), true)
	return !ok
}

// SetIfNotExist is for setting key to hashmap,if exist,return false
func SetIfNotExist(arr []string) bool {
	_, ok := Hashmap.mapping.LoadOrStore(strings.Join(arr, ""), true)
	return !ok
}
