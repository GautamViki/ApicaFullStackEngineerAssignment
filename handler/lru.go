package handler

import (
	dto "apica/dto"
	h "apica/helper"
	"container/list"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	APISuccessCode = "0"
)

// NewLRUCache creates a new LRUCache with the given capacity.
func NewLRUCache(capacity int) *mysqlLru {
	return &mysqlLru{
		lruCache: dto.LRUCache{
			Capacity: capacity,
			Cache:    make(map[int]*list.Element),
			List:     list.New(),
		},
		lru:   dto.Lru{},
		entry: dto.Entry{},
	}
}

type mysqlLru struct {
	lruCache dto.LRUCache
	lru      dto.Lru
	entry    dto.Entry
}

// Get retrieves a value from the cache.
func (l *mysqlLru) Get(w http.ResponseWriter, r *http.Request) {
	urlKey := chi.URLParam(r, "key")
	key, err := strconv.Atoi(urlKey)
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Key is invalid.", "LRU1000")
		return
	}
	var value int
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.MoveToFront(ele)
		value = ele.Value.(*dto.Entry).Value
	}
	res := h.PrepareResponse(APISuccessCode, "Key Fetched Successfully.")
	lruResp := dto.LruResponse{Response: res, LRU: dto.Lru{
		Key:   key,
		Value: value,
	}}
	h.RespondwithJSON(w, http.StatusOK, lruResp)
	return
}

// Set adds a value to the cache.
func (l *mysqlLru) Set(key int, value int) {
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.MoveToFront(ele)
		ele.Value.(*dto.Entry).Value = value
		return
	}
	if l.lruCache.List.Len() == l.lruCache.Capacity {
		l.removeOldest()
	}
	ele := l.lruCache.List.PushFront(&dto.Entry{Key: key, Value: value})
	l.lruCache.Cache[key] = ele
}

// Delete removes a value from the cache.
func (l *mysqlLru) Delete(key int) {
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.Remove(ele)
		delete(l.lruCache.Cache, key)
	}
}

// removeOldest removes the oldest element from the cache.
func (l *mysqlLru) removeOldest() {
	ele := l.lruCache.List.Back()
	if ele != nil {
		l.lruCache.List.Remove(ele)
		delete(l.lruCache.Cache, ele.Value.(*dto.Entry).Key)
	}
}
