function findErrorNums(nums: number[]): number[] {
    const n = nums.length;
    let xor = 0;
    for (const num of nums) {
        xor ^= num;
    }
    for (let i = 1; i <= n; i++) {
        xor ^= i;
    }
    const lowbit = xor & (-xor);
    let num1 = 0, num2 = 0;
    for (const num of nums) {
        if ((num & lowbit) === 0) {
            num1 ^= num;
        } else {
            num2 ^= num;
        }
    }
    for (let i = 1; i <= n; i++) {
        if ((i & lowbit) === 0) {
            num1 ^= i;
        } else {
            num2 ^= i;
        }
    }
    for (const num of nums) {
        if (num === num1) {
            return [num1, num2];
        }
    }
    return [num2, num1];
};