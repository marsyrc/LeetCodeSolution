package array

import "strconv"

func findMissingRanges(nums []int, lower int, upper int) []string {
	res := []string{}
	l := 0

	//special case 1: 数组所有元素完全不再区间内
	if len(nums) == 0 || lower > nums[len(nums)-1] || upper < nums[0] {
		return []string{geneStr(lower, upper)}
	}
	//special case 2 : 处理lower到开头的情况
	if lower < nums[0] {
		res = append(res, geneStr(lower, nums[0]-1))
	}

	i := 0
	for i < len(nums) {
		//遍历到连续区间的末尾
		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		l = i
		if l < len(nums)-1 {
			t := geneStr(nums[l]+1, nums[l+1]-1)
			if t != "" {
				res = append(res, t)
			}
		} else { //到数组尾部，和上界比较
			t := geneStr(nums[l]+1, upper)
			if t != "" {
				res = append(res, t)
			}
		}
		i++
	}
	return res
}

//打印闭区间，不合法输出空串
func geneStr(l, r int) string {
	if l > r {
		return ""
	}
	if l == r {
		return strconv.Itoa(l)
	}
	return strconv.Itoa(l) + "->" + strconv.Itoa(r)
}
