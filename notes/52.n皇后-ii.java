/*
 * @lc app=leetcode.cn id=52 lang=java
 *
 * [52] N皇后 II
 */

// @lc code=start

class Solution {
    List<List<String>> res = new ArrayList<>();
    int ans = 0;

    public int totalNQueens(int n) {
        // 初始化空棋盘
        char[][] board = new char[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                board[i][j] = '.';
            }
        }
        // 回溯 -- 参数为初始棋盘和行row
        backtrace(board, 0);
        return ans;
    }

    void backtrace(char[][] board, int row) {
        // 递归结束
        if (row == board.length) {
            // 将数组转化成List加进去
            ans++;
            return;
        }
        int n = board[row].length;// 列数
        for (int col = 0; col < n; col++) {
            // 排除不合法选择
            if (!isValid(board, row, col)) {
                continue;
            }
            // 做选择
            board[row][col] = 'Q';
            // 递归进入下一层
            backtrace(board, row + 1);
            // 回溯
            board[row][col] = '.';
        }
    }

    // 判断[row][col]是否能放置'Q'
    boolean isValid(char[][] board, int row, int col) {
        // 判断这一列是否产生冲突
        for (int i = 0; i < row; i++) {
            if (board[i][col] == 'Q') {
                return false;
            }
        }
        // 判断左上角是否产生冲突
        for (int i = row - 1, j = col - 1; i > -1 && j > -1; i--, j--) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }
        // 判断右上角是否产生冲突
        for (int i = row - 1, j = col + 1; i > -1 && j < board[i].length; i--, j++) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }
        // 没有冲突返回 true
        return true;
    }

    // 将数组转化成List
    private List<String> convert(char[][] board) {
        List<String> path = new ArrayList<>();
        for (int i = 0; i < board.length; i++) {
            path.add(new String(board[i]));
        }
        return path;
    }
}

// @lc code=end
