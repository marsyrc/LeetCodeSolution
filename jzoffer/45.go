package jianzhioffer

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := []string{}
	for _, n := range nums {
		s := strconv.Itoa(n)
		strs = append(strs, s)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})
	ans := strings.Join(strs, "")
	return ans
}
