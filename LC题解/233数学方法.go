// 使用数学方法 Time: O(log10(n)), Space: O(1)
func countDigitOne(n int) int {
	// 边界情况
	if n < 1 {
	  return 0
	}
	count, factor := 0, 1 // 初始化计数器为0,和位数初始化为1
	for n/factor != 0 {   // 当整数n除以factor不等于0时
	  // 不断执行之下操作
	  // 先求出当前位上的数字digit
	  digit := (n / factor) % 10 // n除以factor再对10取模
	  // 然后计算更高位的数字high
	  high := n/(10*factor) // n除以10倍的factor
	  if digit == 0 {// 如果当前位数的数字等于0
		count+=high*factor // 计数器则加上high乘以factor
	  } else if digit == 1 { // 如果当前位数字等于1
		count+=high*factor // 计数器不仅要加上igh乘以factor
		count+=(n%factor)+1 // 还要加上低位数字(n对factor取模即可求出)再加1
	  } else { // 如果是其他情况
		count+=(high+1)*factor // 计数器则加上(high+1)乘以factor
	  }
	  // 计算完当前factor位上1出现的次数
	  factor = factor*10 // factor乘以10进行更新,准备计算下一个十进制位
	}
	return count // 循环结束后返回count即可。
}
  
//   给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数.
//   示例:
//   输入: 13
//   输出: 6 
//   解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。
