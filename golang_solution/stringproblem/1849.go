package stringproblem

import (
	"strconv"
	"strings"
)

func splitString(s string) bool {
o:
	for i := 1; i < len(s); i++ {
		v, _ := strconv.Atoi(s[:i])
		v--
		for t := s[i:]; t != ""; v-- {
			for len(t) > 1 && t[0] == '0' {
				t = t[1:]
			}
			temp := strconv.Itoa(v)
			if !strings.HasPrefix(t, temp) {
				continue o
			}
			t = t[len(temp):]
		}
		return true
	}
	return false
}
