public class Solution {
    private const int markBase = 10000;
    public bool CircularArrayLoop(int[] nums)
    {
        return Check(nums, a => a > 0) || Check(nums, a => a < 0);
    }

    private bool Check(int[] nums, Predicate<int> predicate)
    {
        int cur, old, mark;
        for (int i = 0; i < nums.Length; i++)
        {
            if (!predicate.Invoke(nums[i]) || nums[i] >= markBase)
            {
                continue;
            }

            cur = i;
            mark = markBase + i;
            while (predicate.Invoke(nums[cur]) && nums[cur] < markBase)
            {
                old = cur;
                cur = (cur + nums[cur]) % nums.Length;

                if (cur < 0)
                    cur += nums.Length;

                if (cur == old)
                    break;

                if (nums[cur] == mark)
                    return true;
                
                nums[old] = mark;
            }
        }

        return false;
    }
}



