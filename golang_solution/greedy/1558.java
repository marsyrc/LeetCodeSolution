package golang_solution.greedy;

class Solution1558 {
    public int minOperations(int[] nums) {
        int res = 0;
        int maxNum = 0;
        for (int num : nums) {
            maxNum = Math.max(maxNum, num);
            while (num != 0) {
                if ((num & 1) != 0)  {
                    res++;
                }
                num >>= 1; 
            }
        }

        if (maxNum != 0) {
            while (maxNum != 0) {
                maxNum >>= 1;
                res++;
            }
            res--;
        }
        return res;
    }
 }