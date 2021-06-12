public class Solution
{
	public string LargestNumber(int[] cost, int target)
	{
		int[,] dp = new int[10, target + 1];
		for (int i = 0; i < 10; ++i)
		{
			for (int j = 0; j <= target; ++j)
			{
				dp[i, j] = int.MinValue;
			}
		}
		int[,] from = new int[10, target + 1];
		dp[0, 0] = 0;
		for (int i = 0; i < 9; ++i)
		{
			int c = cost[i];
			for (int j = 0; j <= target; ++j)
			{
				if (j < c)
				{
					dp[i + 1, j] = dp[i, j];
					from[i + 1, j] = j;
				}
				else
				{
					if (dp[i, j] > dp[i + 1, j - c] + 1)
					{
						dp[i + 1, j] = dp[i, j];
						from[i + 1, j] = j;
					}
					else
					{
						dp[i + 1, j] = dp[i + 1, j - c] + 1;
						from[i + 1, j] = j - c;
					}
				}
			}
		}
		if (dp[9, target] < 0)
		{
			return "0";
		}
		StringBuilder sb = new StringBuilder();
		int curr = 9, next = target;
		while (curr > 0)
		{
			if (next == from[curr, next])
			{
				--curr;
			}
			else
			{
				sb.Append(curr);
				next = from[curr, next];
			}
		}
		return sb.ToString();
	}
}
