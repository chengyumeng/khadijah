package hash

import (
	"strings"
	"sync"
)

type hashmap struct {
	mapping sync.Map
}

var Hashmap hashmap

func (h *hashmap) setIfNotExist(arr []string) bool {
	_, ok := h.mapping.LoadOrStore(strings.Join(arr, ""), true)
	return !ok
}

func SetIfNotExist(arr []string) bool {
	_, ok := Hashmap.mapping.LoadOrStore(strings.Join(arr, ""), true)
	return !ok
}
