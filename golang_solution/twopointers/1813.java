package golang_solution.twopointers;

class Solution1813 {
    public boolean areSentencesSimilar(String sentence1, String sentence2) {
        String[] arr1 = sentence1.split(" ");
        String[] arr2 = sentence2.split(" ");
        int length1 = arr1.length;
        int length2 = arr2.length;
        if (length1 > length2) {
            return areSentencesSimilar(sentence2, sentence1);
        }
        
        int left = 0, right = 0;
        while (left < length1 && left < length2) {
            if (arr1[left].equals(arr2[left])) {
                left++;
            } else {
                break;
            }
        }
        
        if (left == Math.min(length1,length2)) return true;
        //从右边遍历，遇到不同的词停止
        while (right < length1 && right < length2) {
            if (arr1[length1-right-1].equals(arr2[length2 - right -1])) {
                right++;
            } else {
                break;
            }
        }
        return right + left == length1 || right == length1;
    }
}
