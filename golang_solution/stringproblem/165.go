package stringproblem

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return -1 * compareVersion(version2, version1)
	}
	for i := 0; i < n1; i++ {
		num1, _ := strconv.Atoi(strings.TrimLeft(s1[i], " "))
		num2, _ := strconv.Atoi(strings.TrimLeft(s2[i], " "))
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	if n2 > n1 {
		for i := n1; i < n2; i++ {
			cur, _ := strconv.Atoi(strings.TrimLeft(s2[i], " "))
			if cur != 0 {
				return -1
			}
		}
	}
	return 0
}
