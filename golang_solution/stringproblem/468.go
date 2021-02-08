package stringproblem

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if len(IP) < 5 {
		return "Neither"
	}
	if validIPv6(IP) {
		return "IPv6"
	} else if validIPv4(IP) {
		return "IPv4"
	}
	return "Neither"
}

func validIPv4(ip string) bool {
	segs := strings.Split(ip, ".")
	if len(segs) != 4 {
		return false
	}
	for _, s := range segs {
		if s == "" {
			return false
		}
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		num, ok := strconv.Atoi(s)
		if ok != nil || num > 255 || num < 0 {
			return false
		}
	}
	return true
}

func validIPv6(ip string) bool {
	fmt.Println("6")
	segs := strings.Split(ip, ":")
	if len(segs) != 8 {
		return false
	}
	for _, s := range segs {
		if s == "" {
			return false
		}
		if len(s) > 4 {
			return false
		}
		for _, c := range s {
			if !isValid(byte(c)) {
				return false
			}
		}
	}
	return true
}

func isValid(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}
