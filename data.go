package gocache


import (
	"time"
	"bytes"
	"fmt"
)

type CacheType int

const (
	STRING CacheType=iota
	MAP
	SET
)

type cacheData struct {
	data interface{}
	cacheType CacheType
	updateTime time.Time
}



type Set struct {
	m map[interface{}]bool
}

func New() *Set{
	return &Set{
		m:map[interface{}]bool{},
	}
}

func (s *Set) Add(item interface{}){
	s.m[item]=true
}

func (s *Set) Remove(item interface{}){
	delete(s.m,item)
}

func (s *Set) Clear(){
	s.m=map[interface{}]bool{}
}

func (s *Set) Contains(item interface{}) bool{
	_,ok:=s.m[item]
	return ok
}

func (s *Set) Len() int{
	return len(s.m)
}

func (s *Set) Equals(other *Set) bool {
	if other == nil {
		return false
	}
	if s.Len() != other.Len() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) Elements() []interface{} {
	initialLen := len(s.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range s.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (s *Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range s.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}