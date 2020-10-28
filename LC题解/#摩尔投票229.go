/**
229. 求众数 II

给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。

 

示例 1：

输入：[3,2,3]
输出：[3]

示例 2：

输入：nums = [1]
输出：[1]

示例 3：

输入：[1,1,1,3,3,2,2,2]
输出：[1,2]
*/

type candidate struct {
	nodeNum int
	count   int
}

func majorityElement(nums []int) []int {
	res := make ([]int, 0, len(nums))
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

    //大于三分之一的，最多有两个候选人
    //大于n分之一的，最多有n - 1个
	cand1 := candidate{
		nodeNum: nums[0],
		count:   0,
	}
	cand2 := candidate{
		nodeNum: nums[0],
		count:   0,
	}

    //抵消阶段 
	for _, v := range nums {
		if cand1.nodeNum == v {
			cand1.count++
			continue
		}
		if cand2.nodeNum == v {
			cand2.count++
			continue
		}
        //当前候选是否有空的，空的就换成当前元素
		if cand1.count == 0 {
			cand1.nodeNum = v
			cand1.count++
			continue
		}
		if cand2.count == 0 {
			cand2.nodeNum = v
			cand2.count++
			continue
		}
        //当前没有匹配的，也都还能抵消，所以两个都--
		cand1.count--
		cand2.count--
	}
	
    //计数阶段，检验选出的两个人是否符合要求
	cand1.count = 0
	cand2.count = 0 
	for _, v := range nums {
		if cand1.nodeNum == v {
			cand1.count++
		} else if cand2.nodeNum == v {
			cand2.count++
		}
	}
	
	if cand1.count > len(nums) / 3 {
		res = append(res, cand1.nodeNum)
	}
	if cand2.count > len(nums) / 3 {
		res = append(res, cand2.nodeNum)
	}
	return res
}