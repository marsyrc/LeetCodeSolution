// 260. 只出现一次的数字 III

// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

// 示例 :

// 输入: [1,2,1,3,2,5]
// 输出: [3,5]

// 注意：

//     结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
//     你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

// 异或法
// 解题思路

//     1、任何一个数字异或它自己其结果都等于0
//     2、所有数字异或后，出现偶数次的数字都抵消了，只剩下两个出现一次的数
//     3、假设两个数字为a和b，最后异或的结果为c
//     4、我们这时候，只要知道c对应的二进制数中某一个位为1的位数N，比如十进制7二进制表示形式为111，那么可取N=1/2/3，
然后将c与数组中第N位为1的数挨个进行异或，异或结果就是a，b中一个，然后用c异或其中一个数，就可以求出另外一个数了。

// 通过上述方法为什么就能得到问题的解呢？
//     1）c中第N位为1表示a或b中有一个数的第N位也为1，假设该数为a
//     2）如下图所示，当将c与数组中第N位为1的数进行异或时，也就是将x与a外加上其他第N位为1的出现过偶数次的数进行异或，化简即为x与a异或，结果即为b

func singleNumber(arr []int) []int {
	var result []int
	if arr == nil {
		return result
	}
	xor := 0
	// 计算数组中所有数字异或的结果
	for _, v := range arr {
		xor ^= v
	}
	position := 0
	for i := xor; i&1 == 0; i >>= 1 {
		position++
	}
	// 异或的结果与所有第position位为1的数字异或
	// 结果一定是出现一次的两个数中的一个
	tempXor := xor
	for _, v := range arr {
		if (uint(v)>>uint(position))&1 == 1 {
			xor ^= v
		}
	}
	// 得到另外一个出现一次的数字
	return []int{xor, xor ^ tempXor}
}


