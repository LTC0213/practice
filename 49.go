package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, str := range strs {
		bt := []byte(str)
		sort.Slice(bt, func(i, j int) bool {
			return bt[i] < bt[j]
		})
		s := string(bt)
		hash[s] = append(hash[s], str)
	}

	var res [][]string
	for _, h := range hash {
		res = append(res, h)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
