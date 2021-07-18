package slidingwindow

import "sort"

func maxFrequency(nums []int, k int) int {
    res := 1
    sort.Ints(nums)
    l := 0
    total := 0
    for r := 1; r < len(nums); r++ {
        total += (nums[r] - nums[r - 1]) * (r - l)
        
        for total > k {
            total -= nums[r] - nums[l]
            l++
        }
        
        res = max(res, r - l + 1)
    }
    return res 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}