package LRU

type LinkedNode struct {
	key, val  int
	pre, next *LinkedNode
}

type LRUCache struct {
	m          map[int]*LinkedNode
	cap        int
	head, tail *LinkedNode
}

func Constructor(cap int) LRUCache {
	head := &LinkedNode{0, 0, nil, nil}
	tail := &LinkedNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkedNode), cap, head, tail}
}

func (l *LRUCache) moveToFront(node *LinkedNode) {
	head := l.head
	node.pre.next = node.next
	node.next.pre = node.pre

	node.next = head.next
	head.next.pre = node

	node.pre = head
	head.next = node
}

func (l *LRUCache) Get(key int) int {
	cache := l.m
	if v, exist := cache[key]; exist {
		l.moveToFront(v)
		return v.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	head := l.head
	tail := l.tail
	cache := l.m
	if v, exist := cache[key]; exist {
		v.val = value
		l.moveToFront(v)
	} else {
		v := &LinkedNode{key, value, nil, nil}
		if len(cache) == l.cap {
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}
		v.next = head.next
		v.pre = head
		head.next.pre = v
		head.next = v
		cache[key] = v
	}
}
