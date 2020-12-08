package HashTable

func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i, v := range nums {
        if u, ok := dict[target - v]; ok {
            return []int{u, i}
        } else {
            dict[v] = i
        }
    }
    return []int{}
}