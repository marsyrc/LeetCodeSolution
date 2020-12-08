## 伪代码

## BFS

```
q.push(start) // 初始节点放入队列
step = 0 // 初始化步长为0

while not q.empty(): // 当前节点不为空
    ++step // 扩展一层节点
    size = q.size() // 当前这一层节点的个数
    while size-- > 0: // 当前这一层节点数量大于0
        node = q.pop() // 不断出队
        new_nodes = expand(node) // 扩展下一层的节点
        if goal in new_nodes: return step + 1 // 找到结果step+1返回步长
        q.append(new_nodes) // 如果没有找到结果，把下一层的所有节点加入队列

return NOT_FOUND // 队列全部结束，没有找到结果
```

## Bidirectional BFS

```
s1.insert(start) // 方便使用HashSet
s2.insert(end) // start和end放入两个set
step = 0 // 初始化步长为0

while not (s1.empty() || s2.empty()): // 当两个set都不为空，循环执行
    ++step // 步长+1
    swap*(s1, s2) // 交替从左端扩展和从右端扩展
    s = {} // 定义新的空集合
    for node in s1: // 当前需要扩展的这一层节点进行遍历
        new_nodes = expand(node) // 扩展下一层节点
        if any(new_nodes) in s2: return step + 1 // 新的节点在s2集合中，返回step+1，找到路径
        s.append(new_nodes) // 如果没有找到结果，把下一层的所有节点加入队列
    s1 = s // 临时的集合赋值给s1，即把已经走的路径存储起来

return NOT_FOUND // 队列全部结束，没有找到结果
```

# 127单词接龙

## BFS

```go
// BFS Time Complexity: O(n*26^l), l = len(word), n=|wordList| Space Complexity: O(n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
    dict := make(map[string]bool) // 把word存入字典
    for _, word := range wordList {
        dict[word] = true // 可以利用字典快速添加、删除和查找单词
    }
    if _, ok := dict[endWord]; !ok {
        return 0
    }
    // queue := []string{beginWord} 定义辅助队列
    var queue []string
    queue = append(queue, beginWord)

    l := len(beginWord)
    steps := 0

    for len(queue) > 0 {
        steps++
        size := len(queue)
        for i := size; i > 0; i-- { // 当前层级节点
            s := queue[0] // 原始单词
            queue = queue[1:]
            chs := []rune(s)
            for i := 0; i < l; i++ { // 对单词的每一位进行修改
                ch := chs[i]                  // 对当前单词的一位
                for c := 'a'; c <= 'z'; c++ { // 尝试从a-z
                    if c == ch { // 跳过本身比如hot修改为hot
                        continue
                    }
                    chs[i] = c
                    t := string(chs)
                    if t == endWord { // 找到结果
                        return steps + 1
                    }
                    if _, ok := dict[t]; !ok { // 不在dict中，跳过
                        continue
                    }
                    delete(dict, t)          // 从字典中删除该单词，因为已经访问过，若重复访问路径一定不是最短的
                    queue = append(queue, t) // 将新的单词添加到队列
                }
                chs[i] = ch // 单词的第i位复位，再进行下面的操作
            }
        }
    }
    return 0
}
```

## Bidirectional BFS

```go
// Bidirectional BFS Time Complexity: O(n*26^l/2), l = len(word), n=|wordList| Space Complexity: O(n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
    dict := make(map[string]bool) // 把word存入字典
    for _, word := range wordList {
        dict[word] = true // 可以利用字典快速添加、删除和查找单词
    }
    if _, ok := dict[endWord]; !ok {
        return 0
    }
    q1 := make(map[string]bool)
    q2 := make(map[string]bool)
    q1[beginWord] = true // 头
    q2[endWord] = true   // 尾

    l := len(beginWord)
    steps := 0

    for len(q1) > 0 && len(q2) > 0 { // 当两个集合都不为空，执行
        steps++
        // Always expend the smaller queue first
        if len(q1) > len(q2) {
            q1, q2 = q2, q1
        }

        q := make(map[string]bool) // 临时set
        for k := range q1 {
            chs := []rune(k)
            for i := 0; i < l; i++ {
                ch := chs[i]
                for c := 'a'; c <= 'z'; c++ { // 对每一位从a-z尝试
                    chs[i] = c // 替换字母组成新的单词
                    t := string(chs)
                    if _, ok := q2[t]; ok { // 看新单词是否在s2集合中
                        return steps + 1
                    }
                    if _, ok := dict[t]; !ok { // 看新单词是否在dict中
                        continue // 不在字典就跳出循环
                    }
                    delete(dict, t) // 若在字典中则删除该新的单词，表示已访问过
                    q[t] = true     // 把该单词加入到临时队列中
                }
                chs[i] = ch // 新单词第i位复位，还原成原单词，继续往下操作
            }
        }
        q1 = q // q1修改为新扩展的q
    }
    return 0
}
```

## 126、题解

```go
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    dict := make(map[string]bool) // 把word存入字典
    for _, word := range wordList {
        dict[word] = true // 可以利用字典快速添加、删除和查找单词
    }
    if _, ok := dict[endWord]; !ok {
        return [][]string{} // 结尾单词不在dict中，返回空
    }
    delete(dict, beginWord) // 从dict中删除beginWord，变化成beginWord是浪费步数的
    delete(dict, endWord)   // 从dict中删除endWord
    steps := map[string]int{beginWord: 1} // 定义steps哈希表，记录每个单词扩展到这个单词的最短步数
    parents := make(map[string][]string)  // 记录每个单词可以由哪些单词扩展出来，可以回溯构建path
    queue := []string{beginWord}          // 定义辅助队列
    var res [][]string                    // 初始化结果集
    step := 0                             // 初始化当前步数为0
    found := false                        // 记录是否已经找到解，找到解当前层遍历结束，进入下一层
    for len(queue) > 0 && !found { // 队列不为空且没有找到解
        step++ // 步数+1
        size := len(queue)
        for i := size; i > 0; i-- { // 遍历当前层级单词
            s := queue[0]     // 取出队首单词
            queue = queue[1:] // 即为扩展单词的parent
            chs := []rune(s)  // 当前单词用chs表示
            for i := 0; i < len(beginWord); i++ {
                ch := chs[i] // 对当前单词chs的每一位进行替换a-z
                for c := 'a'; c <= 'z'; c++ {
                    if c == ch {
                        continue
                    }
                    chs[i] = c // 进行'a'-'z'的改变
                    t := string(chs)
                    if t == endWord { // 找到结果单词
                        parents[t] = append(parents[t], s) // 记录当前有解的单词的parent列表
                        found = true                       // 标记找到解
                    } else {
                        // Not a new word, but another transform with the same number of steps
                        if v, ok := steps[t]; ok && step < v {
                            parents[t] = append(parents[t], s)
                        } // 不是一个新词但是可以通过其他路径得到，也加入该单词的parent列表
                    }
                    if _, ok := dict[t]; !ok {
                        continue // 如果这个单词不在dict中，跳过
                    }
                    delete(dict, t)                    // 如果该单词在dict中，删除
                    queue = append(queue, t)           // 把该单词加入到队列
                    steps[t] = steps[s] + 1            // 记录步数
                    parents[t] = append(parents[t], s) // 记录该新单词的parent列表
                }
                chs[i] = ch // 还原单词
            }
        }
    }
    if found { // 如果找到解了
        curr := []string{endWord} // 构建最后的目标单词
        getPaths(endWord, beginWord, parents, curr, &res)
    }
    return res
}
func getPaths(word, beginWord string, parents map[string][]string, curr []string, res *[][]string) {
    if word == beginWord {
        swap(curr)
        temp := make([]string, len(curr))
        copy(temp, curr)
        *res = append(*res, temp)
        swap(curr)
        return
    }
    for _, p := range parents[word] {
        curr = append(curr, p)
        getPaths(p, beginWord, parents, curr, res)
        curr = curr[:len(curr)-1]
    }
}
func swap(arr []string) {
    i, j := 0, len(arr)-1
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i, j = i+1, j-1
    }
}
```

