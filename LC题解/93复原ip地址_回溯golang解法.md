#### [93. 复原IP地址](https://leetcode-cn.com/problems/restore-ip-addresses/)

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

**有效的 IP 地址** 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 `0`），整数之间用 `'.' `分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 **有效的** IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 **无效的** IP 地址。

 

**示例 1：**

```
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
```

**示例 2：**

```
输入：s = "0000"
输出：["0.0.0.0"]
```

**示例 3：**

```
输入：s = "1111"
输出：["1.1.1.1"]
```

**示例 4：**

```
输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
```

**示例 5：**

```
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
```

## 回溯模板解法

```go
var res []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	res = []string{}
	var path []string
    
    //回溯函数
    //输入目前待处理字串，与当前段数标记（从1到4）
	var backtrack func (s string, key int)
	backtrack = func (s string, key int) {
		//终止条件：已经做完4段，s长度为0
		if key == 5 {
			if len(s) == 0 {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}

		//选择列表：当前段选择长度为1到3
        //合法性判断，j应当不大于子串长度
		for j := 1; j <= 3 ; j++ {
            
			if j <= len(s) {
				v, err := strconv.Atoi(s[:j])
				if err == nil && v <= 255 {
                    //做选择
					path = append(path, s[:j])
					str := s[j:]
                    //回溯
					backtrack(str, key+1)
                    //取消选择
					path = path[:len(path) - 1]
				}
                //如果当前j= 1, v == 0, 直接剪枝，排除前导0情况
				if v == 0 {
					break
				}
			}
		}
	}
	backtrack(s,1)
	return res
}
```

