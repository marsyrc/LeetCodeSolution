import java.util.*;

class Solution1190 {
    int i = 0;

    public String reverseParentheses(String s) {
        return recur(s);
    }

    public String recur(String s) {
        StringBuilder sb = new StringBuilder();
        while (i < s.length()) {
            if (s.charAt(i) == '(') {
                i++;
                sb.append(recur(s));
            } else if (s.charAt(i) == ')') {
                return reverse(sb.toString());
            } else {
                sb.append(s.charAt(i));
            }
            i++;
        }
        return sb.toString();
    }

    public static String reverse(String str) {
        char[] chars = str.toCharArray();
        String reverse = "";
        for (int i = chars.length - 1; i >= 0; i--) {
            reverse += chars[i];
        }
        return reverse;
    }
}