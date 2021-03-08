package calculator

import (
	"bytes"
	"strconv"
)

//使用两个栈实现（其中操作符栈为单调递增栈）

//定义operation结构体
type operation struct {
	op       byte //'+','-','*','/'
	priority int  //当前操作符的优先级
}

func calculate(s string) int {
	//使用bytes.Buffer实现高效的字符串拼接
	var b bytes.Buffer
	//处理负数的情况
	//数据中存在"-1+4*3*3*3"所以需要处理负数的情况
	//负数出现的情况为在最开头，或者前一个字符为'('的情况
	//解决方法，在-1的前面加一个0即可
	for i := 0; i < len(s); i++ {
		if (s[i] == '-' || s[i] == '+') && (i == 0 || s[i-1] == '(') {
			b.WriteByte('0')
		}
		b.WriteByte(s[i])
	}
	//此处之所以要多添加一个’+‘号，是为了将结果进行输出
	b.WriteByte('+')
	str := b.String()
	//实现两个栈，一个放数字，一个放操作数
	num_stack := construct()
	op_stack := construct()
	//实现一个hash表，存放不同操作符对应的基本的优先级
	hash := make(map[byte]int, 0)
	hash['+'] = 1
	hash['-'] = 1
	hash['*'] = 2
	hash['/'] = 2
	//定义优先级
	prio := 0
	for i := 0; i < len(str); i++ {
		//是空格直接跳过
		if str[i] == ' ' {
			continue
		}
		//如果是操作符
		if _, ok := hash[str[i]]; ok {
			//当前操作符
			oper := operation{str[i], hash[str[i]] + prio}
			//维持一个单调栈
			for !op_stack.empty() && op_stack.peek().(operation).priority >= oper.priority {
				op := op_stack.pop().(operation).op //操作符
				cur := 0                            //本次计算的结果
				//按照栈的进入方式，依次出栈
				b := num_stack.pop().(int)
				a := num_stack.pop().(int)
				switch op {
				case '+':
					cur = a + b
				case '-':
					cur = a - b
				case '*':
					cur = a * b
				case '/':
					cur = a / b
				}
				num_stack.push(cur) //将结果入数字栈
			}
			op_stack.push(oper) //将操作符入栈
		} else if str[i] == '(' { //是左括号，将优先级提升10
			prio += 10
		} else if str[i] == ')' { //是右括号，将优先级减小10
			prio -= 10
		} else {
			//是数字(数字可能有多位数，即34，将连续数字拼接在一块)
			tmp := string(str[i])
			for str[i+1] >= '0' && str[i+1] <= '9' {
				tmp += string(str[i+1])
				i++
			}
			num, _ := strconv.Atoi(tmp)
			num_stack.push(num)
		}
	}
	return num_stack.pop().(int)
}

//手动用链表实现栈
type Node struct {
	Val  interface{}
	Next *Node
}

type stack struct {
	dummy *Node
}

func construct() *stack {
	return &stack{&Node{}}
}

func (s *stack) push(x interface{}) {
	node := &Node{x, nil}
	node.Next = s.dummy.Next
	s.dummy.Next = node
}

func (s *stack) pop() interface{} {
	val := s.dummy.Next.Val
	s.dummy.Next = s.dummy.Next.Next
	return val
}

func (s *stack) peek() interface{} {
	return s.dummy.Next.Val
}

func (s *stack) empty() bool {
	return s.dummy.Next == nil
}
