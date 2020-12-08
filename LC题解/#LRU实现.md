LRU缓存机制（146）

![https://pic.leetcode-cn.com/c8bcb56411f549a2d4c409f021083567e52d1a64748ada00a40ba658bd692d69-image.png](https://pic.leetcode-cn.com/c8bcb56411f549a2d4c409f021083567e52d1a64748ada00a40ba658bd692d69-image.png)


    今天为大家分享很出名的 LRU 算法，第一讲共包括 4 节。
    
        LRU概述
        LRU使用
        LRU实现
        Redis近LRU概述

# LRU 概述

LRU 是 Least Recently Used 的缩写，译为最近最少使用。它的理论基础为“最近使用的数据会在未来一段时期内仍然被使用，已经很久没有使用的数据大概率在未来很长一段时间仍然不会被使用”由于该思想非常契合业务场景 ，并且可以解决很多实际开发中的问题，所以我们经常通过 LRU 的思想来作缓存，一般也将其称为LRU缓存机制。

全网最简单的版本，希望可以坚持看下去！）下面，我们一起学习一下。
第146题：LRU缓存机制
运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作：获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1 。


写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。


进阶:你是否可以在 O(1) 时间复杂度内完成这两种操作？

# Go实现

利用map + 双向链表实现

```go
type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

//func (this *LRUCache) Get(key int) int {
//	head := this.head
//	cache := this.m
//	if v, exist := cache[key]; exist {
//		v.pre.next = v.next
//		v.next.pre = v.pre
//
//		v.next = head.next
//		head.next.pre = v
//
//		v.pre = head
//		head.next = v
//		return v.val
//	}else {
//		return -1
//	}
//}
func (this *LRUCache) moveToHead(node *LinkNode) {
	head := this.head
	//从当前位置删除
	node.pre.next = node.next
	node.next.pre = node.pre
	//移动到首位置
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, exist := cache[key]; exist {
		this.moveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m
	if v, exist := cache[key]; exist {
		v.val = value
		this.moveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == this.cap {
			delete(cache , tail.pre.key)
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
```

