package jianzhioffer

package mathproblem

/*
	输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
*/
func countDigitOne(n int) int {
	//假设某数字形如xyzdabc
	//统计d这一位上1出现的次数
	cnt := 0
	//k相当于是个十进制掩码
	for k := 1; k <= n; k *= 10 {
		abc := n % k
		xyzd := n / k
		d := xyzd % 10
		xyz := xyzd / 10
		//xyz取[0, xyz - 1], abc取[0, k -1]
		cnt += xyz * k

		//xyz取xyz的情况
		//若d==1, 小于等于n的数只有[0, abc], 即abc + 1个
		if d == 1 {
			cnt += abc + 1
		}
		//d > 1, 有k个
		if d > 1 {
			cnt += k
		}
		if xyz == 0 {
			break
		}
	}
	return cnt
}
