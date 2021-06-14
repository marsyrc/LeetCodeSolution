public class Solution {
    public bool IsCovered(int[][] ranges, int left, int right) 
    {
        int[] diff = new int[52];
        foreach (int[] range in ranges) 
        {
            diff[range[0]]++;
            diff[range[1] + 1]--;
        }
        
        int sum = 0;
        for (int i = 1; i <= 50; i++) 
        {
            sum += diff[i];
            if (sum <= 0 && i >= left && i <= right) 
                return false;
        }
        return true;
    }
}