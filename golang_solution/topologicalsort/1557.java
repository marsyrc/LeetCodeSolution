package golang_solution.topologicalsort;

import java.util.*;

class Solution {
    public List<Integer> findSmallestSetOfVertices(int n, List<List<Integer>> edges) {
        int[] indegrees = new int[n];
        for (List<Integer> edge : edges) {
            indegrees[edge.get(1)]++;
        }
        List<Integer> res = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            if (indegrees[i] == 0) {
                res.add(i);
            }
        }
        return res;
    }
}