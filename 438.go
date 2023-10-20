package main

import "fmt"

func findAnagrams(s string, p string) []int {
	hash := make(map[byte]int)
	for i := range p {
		hash[p[i]]++
	}
	count := len(hash)
	m, n := len(s), len(p)
	var res []int
	if m < n {
		return res
	}

	cnt := 0
	for i := 0; i < n; i++ {
		hash[s[i]]--
		if hash[s[i]] == 0 {
			cnt++
		}
	}
	if cnt == count {
		res = append(res, 0)
	}
	for i := n; i < m; i++ {
		hash[s[i]]--
		if hash[s[i]] == 0 {
			cnt++
		}
		if hash[s[i-n]] == 0 {
			cnt--
		}
		hash[s[i-n]]++
		if cnt == count {
			res = append(res, i-n+1)
		}
	}
	return res
}

func main() {
	s, p := "cbaebabacd", "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
