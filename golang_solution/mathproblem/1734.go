package mathproblem

// 给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个 奇数 。
// 它被加密成另一个长度为 n - 1 的整数数组 encoded ，
// 满足 encoded[i] = perm[i] XOR perm[i + 1] 。
// 比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。
// 给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。
func decode(encoded []int) []int {
	res := make([]int, len(encoded)+1)
	n := len(encoded)
	//所有数字的异或和
	xorSum := 0
	for i := 1; i <= n+1; i++ {
		xorSum ^= i
	}
	//除了第一个数字的异或和
	xorSumButFirst := 0
	for i := 1; i < n; i += 2 {
		xorSumButFirst ^= encoded[i]
	}
	res[0] = xorSumButFirst ^ xorSum
	for i := 1; i <= n; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}
