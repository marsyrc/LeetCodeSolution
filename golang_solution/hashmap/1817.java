package golang_solution.hashmap;

import java.util.*;

class Solution1817 {
    public int[] findingUsersActiveMinutes(int[][] logs, int k) {
        int[] ans = new int[k];
        HashMap<Integer, Set<Integer>> cnt = new HashMap<>();
        for (int[] log : logs) {
            int user = log[0], time = log[1];
            Set<Integer> userSet = cnt.getOrDefault(user, new HashSet<>());
            userSet.add(time);
            cnt.put(user, userSet);
        }
        for (Map.Entry<Integer, Set<Integer>> entry : cnt.entrySet()) {
            int n = entry.getValue().size();
            if (n > 0)
                ans[n - 1]++;
        }
        return ans;
    }
}