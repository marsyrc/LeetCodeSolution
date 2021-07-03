function largestArea(grid: string[]): number {
    let m = grid.length, n = grid[0].length;
    let visited = Array.from({ length: m }, v => new Array(n).fill(false));
    let ans = 0;
    let dfs = function (i: number, j: number, target: string): number {
        // 1. 临边岛屿不合法
        if (i < 0 || i > m - 1 || j < 0 || j > n - 1) {
            return Infinity;
        }

        let cur = grid[i].charAt(j);
        // 2. 与0相邻岛屿不合法
        if (cur == '0') {
            return Infinity;
        }

        if (cur != target || visited[i][j]) {
            return 0;
        }

        let ans = 1;
        visited[i][j] = true;
        for (let [dx, dy] of [[0, 1], [0, -1], [1, 0], [-1, 0]]) {
            let x = i + dx, y = j + dy;
            ans += dfs(x, y, target);
        }
        return ans;
    }
    for (let i = 0; i < m; ++i) {
        for (let j = 0; j < n; ++j) {
            if (grid[i].charAt(j) != '0' && !visited[i][j]) {
                let area = dfs(i, j, grid[i].charAt(j));
                ans = Math.max(area == Infinity ? 0 : area, ans);
            }
        }
    }
    return ans;
};
