public class Solution
{
	public bool ConfusingNumber(int n)
	{
		int res = 0;
		int num = n;
		while (n > 0)
		{
			int bit = n % 10;
			n /= 10;

			switch (bit)
			{
				case 2:
				case 3:
				case 4:
				case 5:
				case 7:
					return false;
				case 6:
					res = 10 * res + 9;
					break;
				case 9:
					res = 10 * res + 6;
					break;
				default:
					res = 10 * res + bit;
					break;
			}
		}
		return num > 0 && res != num;
	}
}