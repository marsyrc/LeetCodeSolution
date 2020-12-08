// 384. 打乱数组

// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
// 实现 Solution class:

//     Solution(int[] nums) 使用整数数组 nums 初始化对象
//     int[] reset() 重设数组到它的初始状态并返回
//     int[] shuffle() 返回数组随机打乱后的结果
// 示例：
// 输入
// ["Solution", "shuffle", "reset", "shuffle"]
// [[[1, 2, 3]], [], [], []]
// 输出
// [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
// 解释
// Solution solution = new Solution([1, 2, 3]);
// solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
// solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
// solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]

// 随机洗牌的随机在于其不确定性，对于n张牌组成的有序排列，经过了n次随机选择，漏掉1只牌从未选过的概率不等于0，而且，随着牌的张数数量增加，这个概率非常可观。
// 现在就是经典的Fisher–Yates算法登场的时候了。下面给出伪代码：
//     for(int i=suit.length-1;i>0;i--)
//     {
//         random1 = Random.next(1,i);
//         exchange(suit[random1],suit[i]);
//     }

// 这个算法和之前的算法最大的不同在于，每一次抽卡的范围在慢慢变小

type Solution struct {
	origin []int
	r      *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
		r:      rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.origin)
	res := make([]int, n)
	copy(res, this.origin)
	fmt.Println(res)

	for i := 0; i < n; i++ {
		idx := n - i
		nextRange := this.r.Intn(idx)
		b := i + nextRange
		res[i], res[b] = res[b], res[i]
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
