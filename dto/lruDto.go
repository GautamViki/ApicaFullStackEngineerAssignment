package dto

import (
	"apica/helper"
	"container/list"
)

type LruResponse struct {
	helper.Response
	LRU Lru
}
type Lru struct {
	Key   int
	Value int
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*list.Element
	List     *list.List
}

type Entry struct {
	Key   int
	Value int
}