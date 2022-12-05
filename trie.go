package trie

import "errors"

// ------------------------------------------------- PathSplitFunc -----------------------------------------------------

// PathSplitFunc 用于把传入的路径字符串切割为字典中的一个项，默认是按照字符来切割，使用者可根据自己的需求自定义切割方式
type PathSplitFunc func(s string) ([]string, error)

func DefaultPathSplitFunc(s string) ([]string, error) {
	// 假设大多数情况下是没问题的，可以避免扩容的
	result := make([]string, len(s))
	for index, runeValue := range s {
		if index < len(result) {
			result[index] = string(runeValue)
		} else {
			result = append(result, string(runeValue))
		}
	}
	return result, nil
}

// ------------------------------------------------- Trie --------------------------------------------------------------

type Trie[T any] struct {
	root          *TrieNode[T]
	pathSplitFunc PathSplitFunc
}

func NewTrie[T any](pathSplitFunc ...PathSplitFunc) *Trie[T] {
	return &Trie[T]{
		pathSplitFunc: append(pathSplitFunc, DefaultPathSplitFunc)[0],
		root:          NewTrieNode[T]("", nil),
	}
}

// Add 仅当不存在时插入到树上，已经存在的话则忽略
func (x *Trie[T]) Add(path string, value T) error {
	// 先把路径切割
	pathSlice, err := x.pathSplitFunc(path)
	if err != nil {
		return err
	}
	// 然后遍历路径，把不存在的都补上
	currentNode := x.root
	for index, pathKey := range pathSlice {
		node, exists := currentNode.Children[pathKey]
		if !exists {
			// 如果不存在的话则创建
			node = NewTrieNode(pathKey, currentNode)
			currentNode.Children[pathKey] = node
		} else if index+1 == len(pathSlice) && node.Exists {
			// 如果存在的话则不能覆盖，直接中断返回
			return nil
		}
		currentNode = node
	}
	currentNode.Value = value
	currentNode.Exists = true
	return nil
}

// Upsert 如果树上已经存在则更新，不存在则插入
func (x *Trie[T]) Upsert(path string, value T) error {
	// 先把路径切割
	pathSlice, err := x.pathSplitFunc(path)
	if err != nil {
		return err
	}
	// 然后遍历路径，把不存在的都补上
	currentNode := x.root
	for _, pathKey := range pathSlice {
		node, exists := currentNode.Children[pathKey]
		if !exists {
			node = NewTrieNode(pathKey, currentNode)
			currentNode.Children[pathKey] = node
		}
		currentNode = node
	}
	currentNode.Value = value
	currentNode.Exists = true
	return nil
}

// Remove 从树上移除给定路径
func (x *Trie[T]) Remove(path string) error {
	_, node, err := x.FindTrieNode(path)
	if err != nil {
		return err
	}
	node.Remove()
	return nil
}

// QueryE 查询路径的值
func (x *Trie[T]) QueryE(path string) (value T, err error) {
	_, node, err := x.FindTrieNode(path)
	if err == nil {
		return node.Value, nil
	} else {
		var zero T
		return zero, err
	}
}

// QueryOrDefaultE 查询给定的路径的负载，如果不存在的话则返回默认值
func (x *Trie[T]) QueryOrDefaultE(path string, defaultValue T) (value T, err error) {
	_, node, err := x.FindTrieNode(path)
	if err == nil {
		return node.Value, nil
	}
	if errors.Is(err, ErrNotFound) {
		return defaultValue, nil
	} else {
		var zero T
		return zero, err
	}
}

// FindTrieNode 寻找路径绑定的节点
func (x *Trie[T]) FindTrieNode(path string) ([]string, *TrieNode[T], error) {
	// 先把路径切割
	pathSlice, err := x.pathSplitFunc(path)
	if err != nil {
		return nil, nil, err
	}
	// 根据路径寻找节点
	currentNode := x.root
	for _, pathStep := range pathSlice {
		node, exists := currentNode.Children[pathStep]
		if !exists {
			return nil, nil, ErrNotFound
		}
		currentNode = node
	}
	if currentNode == nil || !currentNode.Exists {
		return nil, nil, ErrNotFound
	}
	return pathSlice, currentNode, nil
}

// ------------------------------------------------- TrieNode ----------------------------------------------------------

type TrieNode[T any] struct {
	Parent   *TrieNode[T]
	Children map[string]*TrieNode[T]
	Exists   bool

	Key   string
	Value T
}

func NewTrieNode[T any](key string, parent *TrieNode[T]) *TrieNode[T] {
	return &TrieNode[T]{
		Parent:   parent,
		Children: make(map[string]*TrieNode[T], 0),
		Exists:   true,

		Key: key,
	}
}

func (x *TrieNode[T]) RemoveChild(key string) {
	delete(x.Children, key)
}

func (x *TrieNode[T]) Remove() {
	// 如果没有孩子节点的话，则直接从树上删除
	if len(x.Children) == 0 {
		if x.Parent != nil {
			x.Parent.RemoveChild(x.Key)
		}
		return
	}
	// 如果有孩子节点的话，将当前节点标记一下，并不将其实际删除
	var zero T
	x.Value = zero
	x.Exists = false
}

// ------------------------------------------------- --------------------------------------------------------------------
