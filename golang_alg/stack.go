package golang_alg

//Stack 栈
type Stack struct {
	data []interface{}
	top  int
	size int
}

//NewStack 初始化栈
func NewStack(len int) *Stack {
	return &Stack{
		data: make([]interface{}, len),
		top:  -1,
		size: len,
	}
}

//IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

//IsNotEmpty 判断栈是否不为空
func (s *Stack) IsNotEmpty() bool {
	return s.top > -1
}

//Size 长度
func (s *Stack) Size() int {
	return s.top + 1
}

//Push 新增元素
func (s *Stack) Push(v interface{}) {
	if s.top+1 == s.size {
		return
	}
	s.top++
	s.data[s.top] = v
}

//Pop 移除栈顶元素
func (s *Stack) Pop() interface{} {
	item := s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return item
}

//Get 获取栈中元素
func (s *Stack) Get(index int) interface{} {
	return s.data[index]
}

//Top 获取栈顶元素
func (s *Stack) Top() interface{} {
	return s.data[s.top]
}
