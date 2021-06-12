public class Solution
{
	int[] prefix;
	int sum;
	Random rd = new Random();
	public Solution(int[] w)
	{
		int n = w.Length;
		prefix = new int[n];
		prefix[0] = w[0];
		for (int i = 1; i < n; i++)
		{
			prefix[i] = prefix[i - 1] + w[i];
		}
		sum = prefix[n - 1];
	}

	public int PickIndex()
	{
		int target = rd.Next(0, sum);

		int left = 0, right = prefix.Length - 1;
		while (left < right)
		{
			int mid = left + (right - left) / 2;
			if (prefix[mid] <= target)
			{
				left = mid + 1;
			}
			else
			{
				right = mid;
			}
		}
		return left;
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(w);
 * int param_1 = obj.PickIndex();
 */