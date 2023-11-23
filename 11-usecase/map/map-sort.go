
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"bb": 20, "aa": 10, "cc": 30}

	//keys := make([]string, 0, len(m))
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
