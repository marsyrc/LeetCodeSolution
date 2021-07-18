package array

func canDivideIntoSubsequences(nums []int, k int) bool {
    if k == 1 {
        return true
    }
    pre := nums[0]
    cnt := 0
    for i := 0; i < len(nums); i++ {
        if pre == nums[i] {
            cnt++
        } else {
            if cnt * k > len(nums) {
                return false
            }
            pre = nums[i]
            cnt = 1
        }
    }
    return cnt * k <= len(nums)
}


// func canDivideIntoSubsequences(nums []int, k int) bool {
//     cnt := make([]int, nums[len(nums) - 1] + 1)
//     for _, num := range nums {
//         cnt[num]++
//     }
//     sort.Ints(cnt)
//     maxC := cnt[len(cnt) - 1]
//     return maxC * k <= len(nums)
// }