package golang_solution.binarysearch;

//交互题
interface FontInfo {
    // Return the width of char ch when fontSize is used.
    public int getWidth(int fontSize, char ch);

    // Return Height of any char when fontSize is used.
    public int getHeight(int fontSize);
}

class Solution1618{
    public int maxFont(String text, int w, int h, int[] fonts, FontInfo fontInfo) {
        int n = fonts.length;
        int l = 0, r = n - 1;

        char[] textArr = text.toCharArray();
        while (l < r) {
            int mid = l + (r - l) / 2 + 1;
            int curWidth = 0, curHeight = fontInfo.getHeight(fonts[mid]);
            for (char c : textArr) {
                curWidth += fontInfo.getWidth(fonts[mid], c);
            }
            if (curWidth > w || curHeight > h) {
                r = mid - 1;
            } else {
                l = mid;
            }
        }
        int nW = 0, nH = fontInfo.getHeight(fonts[l]);
        for (char c : textArr) {
            nW += fontInfo.getWidth(fonts[l], c);
        }
        if (nW > w || nH > h) {
            return -1;
        }
        return fonts[l];
    }
}