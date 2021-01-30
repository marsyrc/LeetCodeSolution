package binarysearch

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

func search(reader ArrayReader, target int) int {
	//找到左右边界
	if reader.get(0) == target {
		return 0
	}
	l, r := 0, 1
	for reader.get(r) < target {
		l = r
		r <<= 1
	}
	//二分查找
	for l <= r {
		mid := l + (r-l)/2
		if reader.get(mid) == target {
			return mid
		} else if reader.get(mid) > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
