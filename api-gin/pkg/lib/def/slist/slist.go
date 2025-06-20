package slist

import (
	"container/list"
	"encoding/json"
	"sync"
)

type SList struct {
	mu   sync.RWMutex
	list *list.List
}

func New() *SList {
	return &SList{
		list: list.New(),
	}
}

func NewByList(l *list.List) *SList {
	return &SList{
		list: l,
	}
}

func NewByArr(arr []any) *SList {
	l := list.New()
	for _, elem := range arr {
		l.PushBack(elem)
	}
	return &SList{
		list: l,
	}
}

func NewByJson(jsonStr string) (*SList, error) {
	var arr []any
	if err := json.Unmarshal([]byte(jsonStr), &arr); err != nil {
		return nil, err
	}

	l := list.New()
	for _, elem := range arr {
		l.PushBack(elem)
	}

	return &SList{
		list: l,
	}, nil
}

// Pushs adds multiple elements to the end of the list.
func (s *SList) Push(v ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, elem := range v {
		s.list.PushBack(elem)
	}
}

// PushFront adds multiple elements to the front of the list.
func (s *SList) PushFront(v ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, elem := range v {
		s.list.PushFront(elem)
	}
}

// Pop removes the last element from the list and returns its value.
func (s *SList) Pop() any {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.list.Len() == 0 {
		return nil
	}
	elem := s.list.Back()
	s.list.Remove(elem)
	return elem.Value
}

// Shift removes the first element from the list and returns its value.
func (s *SList) Shift() any {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.list.Len() == 0 {
		return nil
	}
	elem := s.list.Front()
	s.list.Remove(elem)
	return elem.Value
}

// Len returns the number of elements in the list.
func (s *SList) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.list.Len()
}

// Get returns the value at the specified index.
func (s *SList) Get(idx int) any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if idx < 0 || idx >= s.list.Len() {
		return nil
	}
	elem := s.list.Front()
	for i := 0; i < idx; i++ {
		elem = elem.Next()
	}
	return elem.Value
}

// Set sets the value at the specified index.
func (s *SList) Set(idx int, v any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if idx < 0 || idx >= s.list.Len() {
		return
	}
	elem := s.list.Front()
	for i := 0; i < idx; i++ {
		elem = elem.Next()
	}
	elem.Value = v
}

// Del removes the element at the specified index.
func (s *SList) Del(idx int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if idx < 0 || idx >= s.list.Len() {
		return
	}
	elem := s.list.Front()
	for i := 0; i < idx; i++ {
		elem = elem.Next()
	}
	s.list.Remove(elem)
}

// Clear removes all elements from the list.
func (s *SList) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.list.Init()
}

// Each applies the provided function to each element in the list.
func (s *SList) Each(f func(i int, v any)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, elem := 0, s.list.Front(); elem != nil; i++ {
		f(i, elem.Value)
		elem = elem.Next()
	}
}
