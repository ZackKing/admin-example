package smap

// smap is a safe map implementation for concurrent access.
// It provides a thread-safe way to store and retrieve key-value pairs.

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
)

type SMap struct {
	mu   sync.RWMutex
	data map[any]any
}

func New() *SMap {
	return &SMap{
		data: make(map[any]any),
	}
}

func NewByMap(m map[any]any) *SMap {
	return &SMap{
		data: m,
	}
}

func NewByStruct(d any) *SMap {
	sm := &SMap{
		data: make(map[any]any),
	}

	if d == nil {
		return sm
	}

	v := reflect.ValueOf(d)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return sm
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}
		if idx := strings.Index(key, ","); idx != -1 {
			key = key[:idx]
		}
		if key == "-" {
			continue
		}
		sm.data[key] = v.Field(i).Interface()
	}

	return sm
}

func NewByJson(jsonStr string) (*SMap, error) {
	m := &SMap{
		data: make(map[any]any),
	}
	if err := json.Unmarshal([]byte(jsonStr), &m.data); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *SMap) Set(k string, v any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}

func (m *SMap) Get(k string) (any, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[k]
	return v, ok
}

func (m *SMap) Del(k string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, k)
}

func (m *SMap) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[any]any)
}

func (m *SMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

func (m *SMap) Keys() []any {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]any, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *SMap) Vals() []any {
	m.mu.RLock()
	defer m.mu.RUnlock()
	vals := make([]any, 0, len(m.data))
	for _, v := range m.data {
		vals = append(vals, v)
	}
	return vals
}

func (m *SMap) Each(f func(k, v any)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		f(k, v)
	}
}

func (m *SMap) Json() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	sd := make(map[string]any, len(m.data))
	for k := range m.data {
		if _, ok := k.(string); !ok {
			return ""
		}
		sd[k.(string)] = m.data[k]
	}
	b, _ := json.Marshal(sd)
	return string(b)
}

func (m *SMap) FilterEmpty() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if v == nil || v == "" || v == 0 || v == false {
			delete(m.data, k)
		}
	}
}

func (m *SMap) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if v == nil {
			delete(m.data, k)
		}
	}
}
