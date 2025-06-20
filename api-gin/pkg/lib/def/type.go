package def

import (
	"encoding/json"
	"fmt"
	"sort"
)

type KV map[string]any
type AKV map[any]any

func (a KV) Get(k string, defVal any) any {
	val, ok := a[k]
	if !ok {
		return defVal
	}
	return val
}

func (a KV) Del(k string) {
	delete(a, k)
}

func (a KV) Json() string {
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(b)
}

func GenTree(list []KV, key, pkey, skey string) []KV {
	if len(list) == 0 {
		return []KV{}
	}

	if key == "" {
		key = "id"
	}
	if pkey == "" {
		pkey = "pid"
	}
	if skey == "" {
		skey = "sub"
	}

	items := make(map[any]KV)
	for _, v := range list {
		if id, ok := v[key]; ok {
			items[id] = v
		} else {
			fmt.Printf("Warning: item missing '%s' field, skipping: %v\n", key, v)
		}
	}

	tree := make([]KV, 0)
	for _, item := range items {
		if hasCycle(item, items, key, pkey, make(map[any]bool)) {
			fmt.Printf("Warning: cycle detected for item with id %v, skipping\n", item[key])
			continue
		}

		pid, hasPid := item[pkey]
		if !hasPid || pid == nil || pid == "" || pid == 0 || pid == "0" || (hasPid && fmt.Sprintf("%v", pid) == "0") {
			tree = append(tree, item)
			continue
		}

		if parent, exists := items[pid]; exists {
			if parent[skey] == nil {
				parent[skey] = make([]KV, 0)
			}

			children, ok := parent[skey].([]KV)
			if !ok {
				fmt.Printf("Warning: invalid '%s' field type in parent %v, skipping\n", skey, parent[key])
				continue
			}

			parent[skey] = append(children, item)

			if _, hasSort := item["sort"]; hasSort {
				sort.Slice(parent[skey].([]KV), func(i, j int) bool {
					a := parent[skey].([]KV)[i]["sort"]
					b := parent[skey].([]KV)[j]["sort"]
					aVal, aOk := a.(float64)
					bVal, bOk := b.(float64)
					if aOk && bOk {
						return aVal < bVal
					}

					return fmt.Sprintf("%v", a) < fmt.Sprintf("%v", b)
				})
			}
		} else {
			fmt.Printf("Warning: parent with pid %v not found for item %v\n", pid, item[key])
		}
	}

	if len(tree) > 0 {
		sort.Slice(tree, func(i, j int) bool {
			a := tree[i]["sort"]
			b := tree[j]["sort"]
			aVal, aOk := a.(float64)
			bVal, bOk := b.(float64)
			if aOk && bOk {
				return aVal < bVal
			}
			return fmt.Sprintf("%v", a) < fmt.Sprintf("%v", b)
		})
	}

	return tree
}

func hasCycle(item KV, items map[any]KV, key, pkey string, visited map[any]bool) bool {
	id, ok := item[key]
	if !ok {
		return false
	}
	if visited[id] {
		return true
	}
	visited[id] = true
	defer delete(visited, id)

	pid, ok := item[pkey]
	if !ok {
		return false
	}
	if parent, exists := items[pid]; exists {
		return hasCycle(parent, items, key, pkey, visited)
	}
	return false
}
