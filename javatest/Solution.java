package javatest;

import java.net.Inet4Address;
import java.util.*;

class Solution {
    class Task {
        int id;
        int enqueueTime;
        int processingTime;

        public Task(int id, int enqueueTime, int processingTime) {
            this.id = id;
            this.enqueueTime = enqueueTime;
            this.processingTime = processingTime;
        }
    }

    public int[] getOrder(int[][] tasks) {
        int len = tasks.length;

        List<Task> taskList = new ArrayList<>();
        for (int i = 0; i < len; i++) {
            taskList.add(new Task(i, tasks[i][0], tasks[i][1]));
        }
        Collections.sort(taskList, (t1, t2) -> t1.enqueueTime - t2.enqueueTime);

        PriorityQueue<Task> pq = new PriorityQueue<>((t1, t2) -> {
            if (t1.processingTime == t2.processingTime) {
                return t1.id - t2.id;
            } else {
                return t1.processingTime - t2.processingTime;
            }
        });

        long now = 0;
        int i = 0;
        int[] ret = new int[len];
        int p = 0;

        while (i < len) {
            while (i < len && taskList.get(i).enqueueTime <= now) {
                pq.offer(taskList.get(i));
                i++;
            }
            if (pq.isEmpty()) {
                now = (long) taskList.get(i).enqueueTime;
                while (i < len && taskList.get(i).enqueueTime <= now) {
                    pq.offer(taskList.get(i));
                    i++;
                }
            }
            Task task = pq.poll();
            ret[p++] = task.id;
            now += task.processingTime;
        }
        while (!pq.isEmpty()) {
            ret[p++] = pq.poll().id;
        }
        return ret;
    }

    public static void main(String[] args) {

    }
}

class pair {
    public int first = 0;
    public int second = 0;

    public pair(int a, int b) {
        this.first = a;
        this.second = b;
    }
}