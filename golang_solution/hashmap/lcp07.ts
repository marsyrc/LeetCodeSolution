function numWays(n: number, relation: number[][], k: number): number {


    const map = new Map()
    for (const [s, e] of relation) {
        map.set(s, (map.get(s) || new Set()).add(e))
    }
    let result = 0
    const dfs = (end: Set<Number>, c) => {
        if (c < 0 || !end) return
        if (c === 0 && end.has(n - 1)) return result++

        end.forEach(e => {
            dfs(map.get(e), c - 1)
        })
    }

    dfs(map.get(0), k - 1)
    return result
};