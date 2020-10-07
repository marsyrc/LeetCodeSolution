## 198、easy 

简单dp，可压缩状态转移占用空间。

```go
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	//boundary
	dp1 := nums[0]
	dp2 := max(nums[0], nums[1])

	//状态转移
	for i := 2; i < n; i++ {
		tmp := dp2
		dp2 = max(dp1+nums[i], dp2)
		dp1 = tmp
	}
	return dp2
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

## 213、medium

貌似环形数组+ dp

