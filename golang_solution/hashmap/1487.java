package golang_solution.hashmap;

import java.util.*;

class Solution1487 {
    public String[] getFolderNames(String[] names) {
        int n = names.length;
        String[] res = new String[n];
        Map<String, Integer> hm = new HashMap<>();
        for (int i = 0; i < n; i++) {
            if (hm.containsKey(names[i])) {
                for (int k = hm.get(names[i]) + 1; ; k++) {
                    String newName = names[i] + "(" + Integer.toString(k) + ")";
                    if (!hm.containsKey(newName)) {
                        res[i] = newName;
                        hm.put(names[i], k);
                        hm.put(newName, 0);
                        break;
                    }
                }
            } else {
                res[i] = names[i];
                hm.put(names[i], 0);
            }
        }
        return res;
    }
}