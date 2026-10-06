package hashmap

import (
	"fmt"
)

func IsAnagram(s string, t string) bool {
	fmt.Println(s, t)

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[string]int)
	tMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		sMap[string(s[i])]++
		tMap[string(t[i])]++
	}

	for i := 0; i < len(s); i++ {
		if sMap[string(s[i])] != tMap[string(s[i])] {
			return false
		}

	}

	return true
}
