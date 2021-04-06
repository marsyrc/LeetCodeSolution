package rabinkarp

import "math"

func distinctEchoSubstrings(text string) int {
	const a = 29
	const mod = math.MaxInt64
	n := len(text)
	hash := make([]int64, n+1)
	pow := make([]int64, n+1)

	pow[0] = 1
	for i := 1; i < n; i++ {
		hash[i] = (hash[i-1]*a + int64(text[i-1])) % mod
		pow[i] = pow[i-1] * a % mod
	}

	hm := make(map[int64]bool)
	for i := 0; i < n; i++ {
		for l := 2; i+l-1 < n; l += 2 {
			mid := i + l/2
			hash1 := (hash[mid] - hash[i]*pow[l/2] + mod) % mod
			hash2 := (hash[i+l] - hash[mid]*pow[l/2] + mod) % mod
			if hash1 == hash2 {
				hm[hash1] = true
			}
		}
	}
	return len(hm)
}
