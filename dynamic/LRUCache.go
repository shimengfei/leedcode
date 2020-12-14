package dynamic

type LRUCache struct {
	cache      map[int]*DLinkNode
	capacity   int
	size       int
	head, tail *DLinkNode
}
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		cache:    map[int]*DLinkNode{},
		size:     0,
		capacity: capacity,
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkNode(key, value)
		this.cache[key] = node
		this.addNode(node)
		this.size++
		if this.size > this.capacity {
			rNode := this.removedNode()
			delete(this.cache, rNode.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addNode(node *DLinkNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(rNode *DLinkNode) {
	rNode.prev.next = rNode.next
	rNode.next.prev = rNode.prev
}
func (this *LRUCache) removedNode() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

//每次访问时更新其位置
func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
