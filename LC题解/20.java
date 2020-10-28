/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true

示例 2:

输入: "()[]{}"
输出: true

示例 3:

输入: "(]"
输出: false
 */
//use hashmap Stack 
import java.util.HashMap;
import java.util.Stack;

public class Solution {
    public static void main(String[] args) {
        boolean x = isValid("(){}");
        System.out.print(x);
    }
    public static boolean isValid(String s) {
        HashMap<Character, Character> mappings= new HashMap<Character, Character>();
        mappings.put(')', '(');
        mappings.put('}', '{');
        mappings.put(']', '[');;
        Stack <Character> stack = new Stack <Character>();
        for (int i=0; i<s.length(); i++){
            char c = s.charAt(i);
            if (mappings.containsKey(c)){
                char topElement = stack.empty() ? '#' : stack.pop();
                if (topElement != mappings.get(c)){
                    return false;
                }
            }
            else {
                stack.push(c);
            }
        }
        return stack.isEmpty();
    }
}
