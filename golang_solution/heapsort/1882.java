package golang_solution.heapsort;

import java.util.*;

class Solution {
    public int[] assignTasks(int[] servers, int[] tasks) {
        Queue<Integer> taskq = new LinkedList<>();
        PriorityQueue<Machine> avail = new PriorityQueue<>((o1, o2) -> (o1.wei == o2.wei ? o1.idx - o2.idx : o1.wei - o2.wei));
        PriorityQueue<Machine> working = new PriorityQueue<>((o1, o2) -> (o1.workTo - o2.workTo));

        for (int i = 0; i < servers.length; i++) {
            avail.offer(new Machine(0, servers[i], i));
        }
        int m = tasks.length;
        int[] ans = new int[m];
        for (int i = 0; i < tasks.length; i++) {
            taskq.add(i);

            while (working.size() > 0 && working.peek().workTo <= i) {
                Machine top = working.poll();
                avail.add(top);
            }

            while (avail.size() > 0) {
                if (taskq.size() == 0) {
                    break;
                }
                int curTask = taskq.poll();
                Machine curServer = avail.poll();
                curServer.workTo = i + tasks[curTask];
                ans[curTask] = curServer.idx;
                working.add(curServer);
            }
        }

        int curTime = tasks.length;
        while (taskq.size() > 0) {
            if (working.size() > 0) {
                curTime = working.peek().workTo;
            }
            while (working.size() > 0 && working.peek().workTo <= curTime) {
                Machine top = working.poll();
                avail.add(top);
            }

            while (avail.size() > 0) {
                if (taskq.size() == 0) {
                    break;
                }
                int curTask = taskq.poll();
                Machine curServer = avail.poll();
                curServer.workTo = curTime + tasks[curTask];
                ans[curTask] = curServer.idx;
                working.add(curServer);
            }
        } 
        return ans;
    }
}

class Machine {
    int workTo;
    int wei;
    int idx;

    public Machine(int w1, int w2, int i) {
        this.workTo = w1; 
        this.wei = w2;
        this.idx = i;
    }
}