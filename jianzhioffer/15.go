package jianzhioffer

func hammingWeight(num uint32) (ans int) {
	for i := 0; i <= 31; i++ {
		if num&(1<<i) != 0 {
			ans++
		}
	}
	return
}
