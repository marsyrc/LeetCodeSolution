package golang_solution.heapsort;

import java.util.*;

//same as 846
class solution1296 {
    public boolean isPossibleDivide(int[] nums, int k) {
        Map<Integer, Integer> cntMap = new HashMap<>();
        Arrays.sort(nums);
        for (int num : nums) {
            cntMap.put(num, cntMap.getOrDefault(num, 0) + 1);
        }

        for (int num : nums) {
            int curCnt = cntMap.get(num);
            if (curCnt == 0) {
                continue;
            }
            cntMap.put(num, curCnt - 1);
            for (int i = 1; i < k; i++) {
                if (!cntMap.containsKey(num + i)) {
                    return false;
                }
                int afterCnt = cntMap.get(num + i);
                if (afterCnt == 0) {
                    return false;
                }
                cntMap.put(num + i, afterCnt - 1);
            }
        }
        return true;
    }
}

