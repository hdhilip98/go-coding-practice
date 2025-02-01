package cache

type node struct {
	next  *node
	prev  *node
	key   string
	value string
}

type LRUCache struct {
	capacity   int
	dictionary map[string]*node
	head       *node
	tail       *node
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity:   capacity,
		dictionary: make(map[string]*node, capacity),
	}
}

func (this *LRUCache) Get(key string) (string, bool) {
	node, found := this.dictionary[key]
	if !found {
		return "", false
	}

	this.deleteNode(node)
	this.addToStart(node)
	return node.value, true
}

func (this *LRUCache) Put(key, value string) {
	valueNode, found := this.dictionary[key]

	if found {
		valueNode.value = value
		this.deleteNode(valueNode)
		this.addToStart(valueNode)
	} else {
		if len(this.dictionary) >= this.capacity {
			delete(this.dictionary, this.tail.key)
			this.deleteNode(this.tail)
		}

		valueNode = &node{
			key:   key,
			value: value,
		}
		this.addToStart(valueNode)
		this.dictionary[key] = valueNode
	}
}

func (this *LRUCache) deleteNode(node *node) {
	if node == this.head {
		this.head = node.next
		node.next = nil
	} else if node == this.tail {
		this.tail = node.prev
		node.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		node.prev = nil
		node.next = nil
	}
}

func (this *LRUCache) addToStart(node *node) {
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.head.prev = node
		node.next = this.head
		this.head = node
		node.prev = nil
	}
}
