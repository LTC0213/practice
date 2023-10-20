package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	hash := make(map[byte]struct{})
	for i, j := 0, 0; j < len(s); j++ {
		for _, ok := hash[s[j]]; ok; {
			delete(hash, s[i])
			i++
			_, ok = hash[s[j]]
		}
		hash[s[j]] = struct{}{}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}

func main() {
	s := "abcabcbb"
	//s := "bbbbbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

}
