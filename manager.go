package gocache

import (
	"time"
	"sync"
)

type manager struct {
	cacheMap map[string]cacheData
	ttlMap  map[string]time.Duration
	sync.RWMutex
}


func (m *manager) setString(key string, value string, ttl time.Time){
}

func (m *manager) setString(key string, value string, ttl time.Time){

}

