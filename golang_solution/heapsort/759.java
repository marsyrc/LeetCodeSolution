package golang_solution.heapsort;

import java.util.*;

class Solution759 {
    public List<Interval> employeeFreeTime(List<List<Interval>> avails) {
        List<Interval> ans = new ArrayList<>();
        PriorityQueue<Job> pq = new PriorityQueue<>(
                (o1, o2) -> (avails.get(o1.employee).get(o1.idx).start - avails.get(o2.employee).get(o2.idx).start));
        int ei = 0, anchor = Integer.MAX_VALUE;

        for (List<Interval> employee : avails) {
            pq.offer(new Job(ei++, 0));
            anchor = Math.min(anchor, employee.get(0).start);
        }

        while (pq.size() > 0) {
            Job job = pq.poll();
            Interval iv = avails.get(job.employee).get(job.idx);
            if (anchor < iv.start) {
                ans.add(new Interval(anchor, iv.start));
            }
            anchor = Math.max(anchor, iv.end);
            if (++job.idx < avails.get(job.employee).size()) {
                pq.offer(job);
            }
        }

        return ans;
    }
}

class Job {
    int employee;
    int idx;

    public Job(int employee, int idx) {
        this.employee = employee;
        this.idx = idx;
    }
}

class Interval {
    public int start;
    public int end;

    public Interval() {
    }

    public Interval(int _start, int _end) {
        start = _start;
        end = _end;
    }
}