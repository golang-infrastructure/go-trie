package trie

import (
	"github.com/golang-infrastructure/go-tuple"
	"sync"
)

type SyncTrie[T any] struct {
	lock sync.RWMutex
	trie *Trie[T]
}

func NewSync[T any](pathSplitFunc ...PathSplitFunc) *SyncTrie[T] {
	return &SyncTrie[T]{
		lock: sync.RWMutex{},
		trie: New[T](pathSplitFunc...),
	}
}

// Add 仅当不存在时插入到树上，已经存在的话则忽略
func (x *SyncTrie[T]) Add(path string, value T) error {
	x.lock.Lock()
	defer x.lock.Unlock()

	return x.trie.Add(path, value)
}

// Upsert 如果树上已经存在则更新，不存在则插入
func (x *SyncTrie[T]) Upsert(path string, value T) error {
	x.lock.Lock()
	defer x.lock.Unlock()

	return x.trie.Upsert(path, value)
}

// Remove 从树上移除给定路径
func (x *SyncTrie[T]) Remove(path string) error {
	x.lock.Lock()
	defer x.lock.Unlock()

	return x.trie.Remove(path)
}

// Query 查询路径的值
func (x *SyncTrie[T]) Query(path string) (value T, err error) {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.Query(path)
}

// QueryOrDefault 查询给定的路径的负载，如果不存在的话则返回默认值
func (x *SyncTrie[T]) QueryOrDefault(path string, defaultValue T) (value T, err error) {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.QueryOrDefault(path, defaultValue)
}

// FindTrieNode 寻找路径绑定的节点
func (x *SyncTrie[T]) FindTrieNode(path string) ([]string, *TrieNode[T], error) {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.FindTrieNode(path)
}

// ToSlice 把字典树转为字典列表
func (x *SyncTrie[T]) ToSlice(delimiter ...string) []*tuple.Tuple2[string, T] {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.ToSlice(delimiter...)
}

// QueryByPrefix 根据前缀查询
func (x *SyncTrie[T]) QueryByPrefix(prefix string, delimiter ...string) []*tuple.Tuple2[string, T] {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.QueryByPrefix(prefix, delimiter...)
}

func (x *SyncTrie[T]) Contains(path string) (bool, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()

	return x.trie.Contains(path)
}
