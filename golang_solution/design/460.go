package design

//双向链表+双向链表

type LFUCache struct {
	hm         map[int]*Bucket
	head, tail *Bucket
	n          int
	cnt        int
}

func Constructor(capacity int) LFUCache {
	l := LFUCache{
		hm:   make(map[int]*Bucket),
		head: newBucket(-1),
		tail: newBucket(-1),
		n:    capacity,
		cnt:  0,
	}
	l.head.r = l.tail
	l.tail.l = l.head
	return l
}

func (lfu *LFUCache) Get(key int) int {
	if cur, ok := lfu.hm[key]; ok {
		target := (*Bucket)(nil)
		if cur.r.idx != cur.idx+1 {
			target = newBucket(cur.idx + 1)
			target.r = cur.r
			target.l = cur
			cur.r.l = target
			cur.r = target
		} else {
			target = cur.r
		}

		remove := cur.remove(key)
		target.put(remove.k, remove.v)
		lfu.hm[key] = target

		lfu.DeleteIfEmpty(cur)

		return remove.v
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.n == 0 {
		return
	}
	if cur, ok := lfu.hm[key]; ok {
		cur.put(key, value)
		lfu.Get(key)
	} else {
		if lfu.cnt == lfu.n {
			cur := lfu.head.r
			clear := cur.clear()
			delete(lfu.hm, clear.k)
			lfu.cnt--

			lfu.DeleteIfEmpty(cur)
		}

		first := (*Bucket)(nil)

		if lfu.head.r.idx != 1 {
			first = newBucket(1)
			first.r = lfu.head.r
			first.l = lfu.head
			lfu.head.r.l = first
			lfu.head.r = first
		} else {
			first = lfu.head.r
		}

		first.put(key, value)

		lfu.hm[key] = first
		lfu.cnt++
	}
}

func (lfu *LFUCache) DeleteIfEmpty(cur *Bucket) {
	if cur.isEmpty() {
		cur.l.r = cur.r
		cur.r.l = cur.l
		cur = nil
	}
}

type Item struct {
	l *Item
	r *Item
	k int
	v int
}

func newItem(_k, _v int) *Item {
	return &Item{nil, nil, _k, _v}
}

type Bucket struct {
	l, r       *Bucket
	idx        int
	head, tail *Item
	hm         map[int]*Item
}

func newBucket(_idx int) *Bucket {
	nb := &Bucket{nil, nil, _idx, newItem(-1, -1), newItem(-1, -1), make(map[int]*Item)}
	nb.head.r = nb.tail
	nb.tail.l = nb.head
	return nb
}

func (b *Bucket) put(key, value int) {
	item := &Item{}
	if _, ok := b.hm[key]; ok {
		item = b.hm[key]
		item.v = value
		item.l.r = item.r
		item.r.l = item.l
	} else {
		item = newItem(key, value)
		b.hm[key] = item
	}
	item.r = b.head.r
	item.l = b.head
	b.head.r.l = item
	b.head.r = item
}

func (b *Bucket) remove(key int) *Item {
	if item, ok := b.hm[key]; ok {
		item.l.r = item.r
		item.r.l = item.l

		delete(b.hm, key)
		return item
	}
	return nil
}

func (b *Bucket) clear() *Item {
	item := b.tail.l
	item.l.r = item.r
	item.r.l = item.l
	delete(b.hm, item.k)
	return item
}

func (b *Bucket) isEmpty() bool {
	return len(b.hm) == 0
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
