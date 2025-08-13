package ordermap

import (
	"fmt"
	"sort"
)

// 有序的遍历map
func init() {
	m := make(map[string]string)

	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(m[name])
	}

}
