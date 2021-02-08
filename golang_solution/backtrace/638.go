package backtrace

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	res := 0
	for i := 0; i < len(needs); i++ {
		res += needs[i] * price[i]
	}
	for _, offer := range special {
		isValid := true
		for j := 0; j < n; j++ {
			if needs[j] < offer[j] {
				isValid = false
			}
		}
		if isValid {
			for j := 0; j < n; j++ {
				needs[j] -= offer[j]
			}
			res = min(res, offer[len(offer)-1]+shoppingOffers(price, special, needs))
			for j := 0; j < n; j++ {
				needs[j] += offer[j]
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
