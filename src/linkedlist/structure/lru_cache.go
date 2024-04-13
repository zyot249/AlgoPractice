package structure

/*
Problem:
Ref:
- https://leetcode.com/problems/lru-cache/description/
- https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU

Design a data structure for Least Recently Used cache:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
If the number of keys exceeds the capacity from this operation, evict the least recently used key.

Constraints:
- capacity: [1, 3000]
- key: [0, 10^4]
- value: [0, 10^5]
- At most 2 * 105 calls will be made to get and put.
- The functions get and put must each run in O(1) average time complexity.
*/

/*
First approach:
In order to get in O(1), I think I should use a hashmap to store
For finding the least recently used key, I use a doubly-linked list.
Each time I put a key, I will add the new key to last of list and remove the head if over capacity
Each time I get value of an existing key, I will get the corresponding node in list (mapping by hashmap) and swap it to the end of list
*/

type DoublyLinkedListNode struct {
	Key  int
	Val  int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

type LRUCache struct {
	Capacity int
	Map      map[int]*DoublyLinkedListNode
	Head     *DoublyLinkedListNode
	Tail     *DoublyLinkedListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*DoublyLinkedListNode),
		Head:     nil,
		Tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.Map[key]
	if !exist {
		return -1
	}

	this.updateLru(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.Map[key]
	if exist {
		this.updateLru(node)
		node.Val = value
		return
	}

	this.insert(key, value)

	if len(this.Map) > this.Capacity {
		this.dropLru()
	}
}

func (this *LRUCache) insert(key int, value int) {
	node := &DoublyLinkedListNode{
		Key: key,
		Val: value,
	}

	this.Map[key] = node
	if this.Head == nil {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.Next = node
		node.Prev = this.Tail
		this.Tail = node
	}
}

func (this *LRUCache) dropLru() {
	lruNode := this.Head
	delete(this.Map, lruNode.Key)
	this.Head = this.Head.Next
	this.Head.Prev = nil
}

func (this *LRUCache) updateLru(node *DoublyLinkedListNode) {
	prev := node.Prev
	next := node.Next
	if next == nil {
		return
	}

	if prev != nil {
		prev.Next = next
		next.Prev = prev
	} else {
		this.Head = next
		next.Prev = nil
	}

	this.Tail.Next = node
	node.Next = nil
	node.Prev = this.Tail
	this.Tail = node
}
