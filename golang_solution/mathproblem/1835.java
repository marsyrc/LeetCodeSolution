package golang_solution.mathproblem;

class Solution1835 {
    public int getXORSum(int[] arr1, int[] arr2) {
        int sum1 = XorSum(arr1), sum2 = XorSum(arr2);
        return sum1 & sum2;
    }

    public int XorSum(int[] arr1) {
        int result = 0;
        for (int i = 0; i < arr1.length; i++) {
            result ^= arr1[i];
        }
        return result;
    }
}