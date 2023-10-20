package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	hash := make(map[byte]int)
	for i := range t {
		hash[t[i]]++
	}
	count := len(hash)
	var res string

	for i, j, cnt := 0, 0, 0; j < len(s); j++ {
		hash[s[j]]--
		if hash[s[j]] == 0 {
			cnt++
		}
		for i <= j && hash[s[i]] < 0 {
			hash[s[i]]++
			i++
		}
		if cnt == count {
			if len(res) == 0 || len(res) > j-i+1 {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}
