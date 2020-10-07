/*
 * @lc app=leetcode.cn id=17 lang=java
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

//标准回溯


class Solution {
    public List<String> letterCombinations(String digits) {
        List<String> combinations = new ArrayList<String>();

        if (digits.length() == 0)
            return combinations;

        Map<Character, String> phoneMap = new HashMap<Character, String>() {
            {
                put('2', "abc");
                put('3', "def");
                put('4', "ghi");
                put('5', "jkl");
                put('6', "mno");
                put('7', "pqrs");
                put('8', "tuv");
                put('9', "wxyz");
            }
        };

        backtrack(combinations, phoneMap, digits, 0, new StringBuffer());

        return combinations;

    }

    public void backtrack(List<String> combinations, Map<Character, String> phoneMap, String digits, int index, StringBuffer combination){
            //满足终止条件
            if (index == digits.length()){
                combinations.add(combination.toString());
            } else{
                char digit = digits.charAt(index);
                String letters = phoneMap.get(digit);
                int lttersCount = letters.length();

                //遍历选择列表
                for(int i = 0 ; i< lttersCount; ++i){

                    //做选择
                    combination.append(letters.charAt(i));

                    backtrack(combinations, phoneMap, digits, index + 1, combination);
                    
                    //撤销
                    combination.deleteCharAt(index);
                }
            }
    }
}
// @lc code=end
