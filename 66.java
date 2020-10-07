/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
示例 1:
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
 */
class Solution{
    public int[] plusone(int[] digits){
        for (int i = digits.length - 1; i >= 0; i--){// 考虑9+1的进位问题，维护移位循环
            digits[i]++;
            digits[i] = digits[i] % 10;
            if (digits[i] != 0)return digits;  //本位不是0,说明没进位.于是跳出,返回digits
        }
        digits = new int[digits.length + 1]; //for结束说明结果是1000000...
        digits[0] = 1;
        return digits;
    }
}