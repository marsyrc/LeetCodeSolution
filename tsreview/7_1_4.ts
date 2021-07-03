/*
    base case 过不了，bs(...,...,...,...,right, nums1, nums2) NaN
*/
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    let n: number = nums1.length;
    let m: number = nums2.length;

    let left: number = (n + m + 1) / 2;
    let right: number = (n + m + 2) / 2;
    return bs(0, n - 1, 0, m - 1, left, nums1, nums2) * 0.5 + bs(0, n - 1, 0, m - 1, right, nums1, nums2) * 0.5
};

 function bs(start1: number, end1: number, start2: number, end2: number, k: number, nums1: number[], nums2: number[]): number {
    let len1 = end1 - start1 + 1;
    let len2 = end2 - start2 + 1;
    if (len1 > len2) {
        return bs(start2, end2, start1, end1, k, nums2, nums1)
    }
    if (len1 == 0) {
        return nums2[start2 + k - 1]
    }
    if (k == 1) {
        return Math.min(nums1[start1], nums2[start2])
    }

    let i = start1 + Math.min(len1, k / 2) - 1;
    let j = start2 + Math.min(len2, k / 2) - 1;

    if (nums1[i] > nums2[j]) {
        return bs(start1, end1, j + 1, end2, k - (j - start2 + 1), nums1, nums2);
    } else {
        return bs(i + 1, end1, start2, end2, k - (i - start1 + 1), nums1, nums2);
    }
}