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
func (l *mysqlLru) GetByKey(w http.ResponseWriter, r *http.Request) {
	urlKey := chi.URLParam(r, "key")
	key, err := strconv.Atoi(urlKey)
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Key is invalid.", "LRU1000")
		return
	}
	var value *int
	defaultValue := -1
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.MoveToFront(ele)
		value = &ele.Value.(*dto.Entry).Value
	} else {
		value = &defaultValue
	}
	res := h.PrepareResponse(APISuccessCode, "Key Fetched Successfully.")
	lruResp := dto.LruResponse{Response: res, LRU: dto.Lru{
		Key:   key,
		Value: value,
	}}
	h.RespondwithJSON(w, http.StatusOK, lruResp)
}

func (l *mysqlLru) GetAll(w http.ResponseWriter, r *http.Request) {
	lrus := []dto.Lru{}
	for key, _ := range l.lruCache.Cache {
		ele := l.lruCache.Cache[key]
		lrus = append(lrus, dto.Lru{Key: key, Value: &ele.Value.(*dto.Entry).Value})
	}
	res := h.PrepareResponse(APISuccessCode, "Cache Fetched Successfully.")
	h.RespondwithJSON(w, http.StatusOK, dto.LrusResponses{Response: res, Lrus: lrus})
}

// Set adds a value to the cache.
func (l *mysqlLru) Set(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key, err := strconv.Atoi(query.Get("key"))
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Key is invalid.", "LRU1001")
		return
	}
	value, err := strconv.Atoi(query.Get("value"))
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Value is invalid.", "LRU1002")
		return
	}
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.MoveToFront(ele)
		ele.Value.(*dto.Entry).Value = value
	}
	if l.lruCache.List.Len() == l.lruCache.Capacity {
		l.removeOldest()
	}
	ele := l.lruCache.List.PushFront(&dto.Entry{Key: key, Value: value})
	l.lruCache.Cache[key] = ele

	res := h.PrepareResponse(APISuccessCode, "Set Value in cache key successfully.")
	h.RespondwithJSON(w, http.StatusCreated, res)
}

// Delete removes a value from the cache.
func (l *mysqlLru) Delete(w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(chi.URLParam(r, "key"))
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Key is invalid.", "LRU1004")
		return
	}
	if ele, ok := l.lruCache.Cache[key]; ok {
		l.lruCache.List.Remove(ele)
		delete(l.lruCache.Cache, key)
	}

	res := h.PrepareResponse(APISuccessCode, "Cache key deleted successfully.")
	h.RespondwithJSON(w, http.StatusOK, res)
}

// removeOldest removes the oldest element from the cache.
func (l *mysqlLru) removeOldest() {
	ele := l.lruCache.List.Back()
	if ele != nil {
		l.lruCache.List.Remove(ele)
		delete(l.lruCache.Cache, ele.Value.(*dto.Entry).Key)
	}
}
