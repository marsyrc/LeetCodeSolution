import java.util.HashMap;
/*给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
。*/
public class Solution{
		public static void main(String[] args) {
			int x = lengthOfLongestSubstring("hudhubhuddhudb");
			System.out.print(x);
		}
		public static int lengthOfLongestSubstring(String s) {
			if (s.length()==0)return 0;
			HashMap<Character,Integer> map = new HashMap <Character, Integer>();
			int max = 0;
			int left = 0;
			for (int i =0; i < s.length(); i++) {
				if (map.containsKey(s.charAt(i))) {
				left = Math.max(left,map.get(s.charAt(i)) +1);
			}
			map.put(s.charAt(i),i);
			max = Math.max(max, i-left+1);
			}
		return max;
		}
};
//队列思想，以abca为例，到i=3时发现重复，于是把i=0移除，left=1;
//如果是abcb,就得把ab移除，left=2;