package mathproblem

//摩尔投票法
func majorityElement(nums []int) (res []int) {
	cand1 := nums[0]
	cand2 := nums[0]
	cnt1, cnt2 := 0, 0
	//vote process
	for _, n := range nums {
		if n == cand1 {
			cnt1++
			continue
		}
		if n == cand2 {
			cnt2++
			continue
		}
		if cnt1 == 0 {
			cand1 = n
			cnt1 = 1
			continue
		}
		if cnt2 == 0 {
			cand2 = n
			cnt2 = 1
			continue
		}
		cnt1--
		cnt2--
	}
	cnt1 = 0
	cnt2 = 0
	for _, n := range nums {
		if n == cand1 {
			cnt1++
		} else if n == cand2 {
			cnt2++
		}
	}
	if cnt1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return
}
