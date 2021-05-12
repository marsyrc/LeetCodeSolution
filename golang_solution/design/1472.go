package design

type BrowserHistory struct {
	ls  [5001]string
	idx int
	top int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		ls:  [5001]string{homepage},
		idx: 0,
		top: 1,
	}
}

func (b *BrowserHistory) Visit(url string) {
	b.idx++
	b.top = b.idx
	b.ls[b.top] = url
	b.top++
}

func (b *BrowserHistory) Back(steps int) string {
	steps = min(b.idx, steps)
	b.idx -= steps
	return b.ls[b.idx]
}

func (b *BrowserHistory) Forward(steps int) string {
	steps = min(steps, b.top-b.idx-1)
	b.idx += steps
	return b.ls[b.idx]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
