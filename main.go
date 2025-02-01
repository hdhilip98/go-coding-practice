package main

import "github.com/hdhilip98/go-coding-practice/cache"

func main() {
	lruCache := cache.Constructor(3)

	lruCache.Put("A", "A")
	lruCache.Put("B", "B")
	lruCache.Put("C", "C")
	lruCache.Get("A")
	lruCache.Put("D", "D")
}
