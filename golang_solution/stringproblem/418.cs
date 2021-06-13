public class Solution {
    public int WordsTyping(string[] sentence, int rows, int cols) {
        int[] dp = new int[sentence.Length];
        int[] next = new int[sentence.Length];
        
        for (int i = 0; i < sentence.Length; i++) 
        {
            int count = 0;
            int ptr = i;
            int cur = cols;
            while (cur >= sentence[ptr].Length) {
                cur -= sentence[ptr].Length + 1;
                ++ptr;
                if (ptr == sentence.Length) {
                    ++count;
                    ptr = 0;
                }
            }
            dp[i] = count;
            next[i] = ptr;
        }
        
        int res = 0;
        int j = 0;
        for (int i = 0; i < rows; i++) {
            res += dp[j];
            j = next[j];
        }
        return res;
    }
}