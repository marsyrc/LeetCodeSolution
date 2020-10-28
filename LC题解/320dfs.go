import "strconv"

// 320. 列举单词的全部缩写

// 请你写出一个能够举单词全部缩写的函数。

// 注意：输出的顺序并不重要。

// 示例：

// 输入: "word"
// 输出:
// ["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]

func generateAbbreviations(word string) []string {
	n := len(word)
	res := []string{}
	var dfs func(index int, cnt int, path string)
	//cnt计数当前未放入字符的数量,index下表，path当前路径
	dfs = func(index int, cnt int, path string) {
		if index == n {
			if cnt != 0 { //如果到底了计数不为0，就加上计数
				path += strconv.Itoa(cnt)
			}
			res = append(res, path)
			return
		}

		//做选择，递归
		if cnt == 0 {
			dfs(index+1, 0, path+string(word[index])) //直接加当前字符
		} else {
			dfs(index+1, 0, path+strconv.Itoa(cnt)+string(word[index])) //加上前面的计数后，加当前字符
		}
		dfs(index+1, cnt+1, path) //不加当前字符，计数++
	}
	dfs(0, 0, "")
	return res
}
