using System;
using System.Collections.Generic;


public class Solution
{
    public class Point
    {
        public long x;
        public long y;
        public Point(long xx, long yy)
        {
            x = xx;
            y = yy;
        }

        // override object.Equals
        public override bool Equals(object obj)
        {
            //
            // See the full list of guidelines at
            //   http://go.microsoft.com/fwlink/?LinkID=85237
            // and also the guidance for operator== at
            //   http://go.microsoft.com/fwlink/?LinkId=85238
            //

            if (obj == null || GetType() != obj.GetType())
                return false;

            var pp = (Point)obj;
            return x == pp.x && y == pp.y;
        }

        // override object.GetHashCode
        public override int GetHashCode()
        {
            return (int)(x ^ y);
        }
    }
    public int MaxPoints(int[][] points)
    {
        int n = points.Length;
        if (n < 3)
        {
            return n;
        }
        int res = 0;
        Point same = new Point(0, 0);

        for (int i = 0; i < n - 1; i++)
        {
            Dictionary<Point, int> hm = new Dictionary<Point, int>();
            int temp = 0;
            for (int j = i + 1; j < n; j++)
            {
                Point k = GetPoint(points[i], points[j]);
                int val = 0;
                if (hm.TryGetValue(k, out val))
                    hm[k] = val + 1;
                else
                    hm.Add(k, 1);
                if (k.x != 0 || k.y != 0)
                    temp = Math.Max(temp, hm[k]);
            }
            int cnt = 0;
            if (hm.ContainsKey(same))
                cnt = hm[same];
            res = Math.Max(res, temp + 1 + cnt);
        }
        return res;
    }

    private long gcd(long a, long b)
    {
        if (a == 0)
            return b;
        return gcd(b % a, a);
    }

    public Point GetPoint(int[] p1, int[] p2)
    {
        long dx = (long)(p1[0] - p2[0]);
        long dy = (long)(p1[1] - p2[1]);
        if (dx == 0 && dy == 0)
            return new Point(0, 0);
        if (dx == 0)
            return new Point(1, 0);
        if (dy == 0)
            return new Point(0, 1);
        if (dx < 0)
        {
            dx = -dx;
            dy = -dy;
        }
        long d = gcd(dx, dy);
        return new Point(dx / d, dy / d);
    }
}