首先，还是一样的思路：如何穷举？这里的穷举思路和上篇文章递归的思想不太一样。

递归其实是符合我们思考的逻辑的，一步步推进，遇到无法解决的就丢给递归，一不小心就做出来了，可读性还很好。缺点就是一旦出错，你也不容易找到错误出现的原因。比如上篇文章的递归解法，肯定还有计算冗余，但确实不容易找到。

而这里，我们不用递归思想进行穷举，而是利用「状态」进行穷举。

看看总共有几种「状态」，再找出每个「状态」对应的「选择」。我们要穷举所有「状态」，穷举的目的是根据对应的「选择」更新状态。看图，就是这个意思。

![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3aeibzourMAl7wGxtQfWTMtHs5QVdLYibrKZ2RgrqujZkLJEnrn7DibNBbg/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

具体到当前问题，**每天都有三种「选择」**：买入、卖出、无操作，我们用 buy, sell, rest 表示这三种选择。

但问题是，并不是每天都可以任意选择这三种选择的，因为 sell 必须在 buy 之后，buy 必须在 sell 之后（第一次除外）。那么 rest 操作还应该分两种状态，一种是 buy 之后的  rest（持有了股票），一种是 sell 之后的 rest（没有持有股票）。而且别忘了，我们还有交易次数 k 的限制，就是说你 buy 还只能在 k > 0 的前提下操作。

很复对吧，不要怕，我们现在的目的只是穷举，你有再多的状态，老夫要做的就是一把梭全部列举出来。**这个问题的「状态」有三个**，第一个是天数，第二个是当天允许交易的最大次数，第三个是当前的持有状态（即之前说的 rest 的状态，我们不妨用 1 表示持有，0 表示没有持有）。

我们用一个三维数组 dp 就可以装下这几种状态的全部组合，用 for 循环就能完成穷举：



![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3aSvQTuRMLrsMKSB8BQteEbiakEP2fJyNjyVUiaZQLjJ8mlQ7dM4hLYfeQ/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)



而且我们可以用自然语言描述出每一个状态的含义，比如说 dp[3][2][1] 的含义就是：今天是第三天，我现在手上持有着股票，至今最多进行 2 次交易。再比如 dp[2][3][0] 的含义：今天是第二天，我现在手上没有持有股票，至今最多进行 3 次交易。很容易理解，对吧？

我们想求的最终答案是 dp[n - 1][K][0]，即最后一天，最多允许 K 次交易，所能获取的最大利润。读者可能问为什么不是 dp[n -  1][K][1]？因为 [1] 代表手上还持有股票，[0] 表示手上的股票已经卖出去了，很显然后者得到的利润一定大于前者。

现在，我们完成了「状态」的穷举，我们开始思考每种「状态」有哪些「选择」，应该如何更新「状态」。



因为我们的选择是 buy, sell, rest，而这些选择是和「持有状态」相关的，所以只看「持有状态」，可以画个状态转移图。



![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3auom48UuPXWBQV7IRsa59yHe1TRQicDquqyRFOoia62BqHzofboyiay5IA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)



通过这个图可以很清楚地看到，每种状态（0 和 1）是如何转移而来的。根据这个图，我们来写一下状态转移方程：



![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3aLFHicK3LhVZXEJvHzEOgGpjp8RzCxIkQpW0K7qGkqYKcCP5jdJIrpibA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)



这个解释应该很清楚了，如果 buy，就要从利润中减去 prices[i]，如果 sell，就要给利润增加  prices[i]。今天的最大利润就是这两种可能选择中较大的那个。而且注意 k 的限制，我们在选择 buy 的时候，把最大交易数 k 减小了  1，很好理解吧，当然你也可以在 sell 的时候减 1，一样的。

现在，我们已经完成了动态规划中最困难的一步：状态转移方程。**如果之前的内容你都可以理解，那么你已经可以秒杀所有问题了，只要套这个框架就行了。**不过还差最后一点点，就是定义 base case，即最简单的情况。



![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3akByqn8e7kyr0hSKS6iaVkicDsZrc08oic4wp5c7sPk7LzicGJm3xlBRSew/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

把上面的状态转移方程总结一下：

![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4nPicwNq5syrSwnBc02yxG3aewN24fa7UR8G7byHOb7lUfrlgCkUN1KsL5PYsIicKfE0mQ2OCibCXiajA/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)



## Leetcode  121\ 122\123\188\309\714

## 188。任意交易次数

```go
func maxProfit(k int, prices []int) int {
	if k>len(prices)/2{
		return maxProfitInf(prices)
	}
	dp_i_0, dp_i_1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i < k+1; i++{
		dp_i_0[i] = 0
		dp_i_1[i] = math.MinInt64
	}
	for _, v := range prices{
		for i := k; i >=1; i--{
			dp_i_0[i] = max(dp_i_0[i], dp_i_1[i] + v)
			dp_i_1[i] = max(dp_i_1[i], dp_i_0[i-1] - v)
		}
	}
	return dp_i_0[k]
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func maxProfitInf(prices []int) int {
	if len(prices)<2{
		return 0
	}
	sum:=0
	plen:=len(prices)
	for i:=1;i<plen;i++{
		if prices[i]-prices[i-1]>0{
			sum+=prices[i]-prices[i-1]
		}
	}
	return sum
}
```

## 714。含手续费

```go
func maxProfit(prices []int, fee int) int {
	dpI0 := 0
	dpI1 := math.MinInt64
	for _, v := range prices {
		tmp := dpI0
		dpI0 = max(dpI0, dpI1+ v)
		dpI1 = max(dpI1, tmp - v - fee)
	}
	return dpI0
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

