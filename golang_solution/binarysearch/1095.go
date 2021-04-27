package binarysearch

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	find := func(left bool) int {
		l, r := 0, mountainArr.length()-1
		for l < r {
			mid := l + (r-l)/2
			cur, next := mountainArr.get(mid), mountainArr.get(mid+1)
			// 山脉上升
			if cur < next {
				// 即使是满足条件也需要指定才向左寻找
				if left && cur >= target {
					r = mid
				} else {
					l = mid + 1
				}
			} else {
				// 即使是满足条件也需要指定才向右寻找
				if cur > target && !left {
					l = mid + 1
				} else {
					r = mid
				}
			}
		}
		return l
	}
	left := find(true)
	if mountainArr.get(left) == target {
		return left
	}
	right := find(false)
	if mountainArr.get(right) == target {
		return right
	}
	return -1
}
