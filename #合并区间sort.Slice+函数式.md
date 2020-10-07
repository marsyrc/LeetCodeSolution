## 按照切片头排序的方式

	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})

golang sort package: [https://golang.org/src/sort](https://link.jianshu.com/?t=https://golang.org/src/sort)

sort 操作的对象通常是一个 slice，需要满足三个基本的接口，并且能够使用整数来索引

```go
// A type, typically a collection, that satisfies sort.Interface can be 
// sorted by the routines in this package. The methods require that the 
// elements of the collection be enumerated by an integer index. 

type Interface interface { 
 // Len is the number of elements in the collection.
 Len() int

 // Less reports whether the element with
 // index i should sort before the element with index j.
 Less(i, j int) bool

 // Swap swaps the elements with indexes i and j.
 Swap(i, j int)
}
```

ex-1 对 []int 从小到大排序

```go
package main

import (
    "fmt"
    "sort"
)
type IntSlice []int
func (s IntSlice) Len() int { return len(s) 

func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
    a := []int{4,3,2,1,5,9,8,7,6}
    sort.Sort(IntSlice(a))
    fmt.Println("After sorted: ", a)
}
```

ex-2 使用 sort.Ints 和 sort.Strings
 golang 对常见的 []int 和 []string  分别定义了 IntSlice 和 StringSlice, 实现了各自的排序接口。而 sort.Ints 和 sort.Strings  可以直接对 []int 和 []string 进行排序, 使用起来非常方便

```go
package main

import (
    "fmt"
    "sort"
)
func main() {
    a := []int{3, 5, 4, -1, 9, 11, -14}
    sort.Ints(a)
    fmt.Println(a)
    ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
    sort.Strings(ss)
    fmt.Println(ss)
    sort.Sort(sort.Reverse(sort.StringSlice(ss)))
    fmt.Printf("After reverse: %v\n", ss)
}



 
```

ex-3 使用 sort.Reverse 进行逆序排序
 如果我们想对一个 sortable object 进行逆序排序，可以自定义一个type。但 sort.Reverse 帮你省掉了这些代码

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    a := []int{4,3,2,1,5,9,8,7,6}
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    fmt.Println("After reversed: ", a)
}
```

ex-4 使用 sort.Stable 进行稳定排序
 sort.Sort 并不保证排序的稳定性。如果有需要, 可以使用 sort.Stable

```go
package main

import (
    "fmt"
    "sort"
)

type person struct {
  Name string
  Age int
}

type personSlice []person

func (s personSlice) Len() int { return len(s) }

func (s personSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
    a := personSlice {
      {
        Name: "AAA",
        Age: 55,
      },
      {
        Name: "BBB",
        Age: 22,
      },
      {
        Name: "CCC",
        Age: 0,
      },
      {
        Name: "DDD",
        Age: 22,
      },
      {
        Name: "EEE",
        Age: 11,
      },  
    }
    sort.Stable(a)
    fmt.Println(a)
}
```

#### [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/)

给出一个区间的集合，请合并所有重叠的区间。

 

**示例 1:**

```
输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
```

**示例 2:**

```
输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
```

**注意：**输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

 

```go
import "sort"

func merge(intervals [][]int) [][]int {
	if intervals == nil {
		return [][]int{}
	}
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	for i := 0; i < len(intervals); i++ {
		if len(ans) == 0 || intervals[i][0] > ans[len(ans)-1][1] {
			ans = append(ans, intervals[i])
		} else if intervals[i][1] > ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = intervals[i][1]
		}
	}
	return ans
}

```

