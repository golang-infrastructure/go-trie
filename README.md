# 字典树的Go实现



# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-trie
```

# 二、示例代码

## 2.1 基本单词查询

```go
package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-trie"
)

func main() {

	trie := trie.New[string]()

	err := trie.Add("china", "中国")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = trie.Add("chinese", "中国人")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := trie.Query("china")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
	// Output:
	// 中国

}
```

## 2.2 Web路由 

```go
package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-trie"
	"strings"
)

func main() {

	// 假装是一个Web路由
	trie := trie.New[func() error](func(s string) ([]string, error) {
		slice := make([]string, 0)
		for _, x := range strings.Split(s, "/") {
			if x != "" {
				slice = append(slice, x)
			}
		}
		return slice, nil
	})

	// 增加一个路由
	err := trie.Add("/foo/bar", func() error {
		fmt.Println("路由到了/foo/bar")
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 尝试寻找不存在的路由
	_, err = trie.Query("/foo")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 尝试寻找存在的路由
	handler, err := trie.Query("//foo//bar/")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = handler()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Output:
	// not found
	// 路由到了/foo/bar

}
```

# 三、API

## PathSplitFunc

用于把传入的路径字符串切割为字典中的一个项，默认是按照字符来切割，使用者可根据自己的需求自定义切割方式

```go
type PathSplitFunc func(s string) ([]string, error)
```

## New

创建一个字典树

```go
func New[T any](pathSplitFunc ...PathSplitFunc) *Trie[T]
```

## Add

往字典树上添加一个单词，单词可以绑定一个值，但是如果要添加单词已经存在的话则添加失败，单词原来绑定的值不会被修改

```go
func (x *Trie[T]) Add(path string, value T) error
```

## Upsert

往字典树上添加一个单词，单词可以绑定一个值，如果单词已经存在的话则更新其绑定的值，如果不存在的话则新增 

```go
func (x *Trie[T]) Upsert(path string, value T) error 
```

## Remove

从字典树上移除单词 

```go
func (x *Trie[T]) Remove(path string) error
```

## Query

查询给定的单词绑定的值，如果给定的单词在字典树上不存在的话，则返回`ErrNotFound`错误

```go
func (x *Trie[T]) Query(path string) (value T, err error)
```

## QueryOrDefault

查询给定的单词绑定的值，如果给定的单词在字典树上不存在的话，则返回给定的默认值 

```go
func (x *Trie[T]) QueryOrDefault(path string, defaultValue T) (value T, err error)
```

## FindTrieNode

根据给定的单词查找其对应的字典树上的节点，一般用不到，都是内部API使用 

```go
func (x *Trie[T]) FindTrieNode(path string) ([]string, *TrieNode[T], error)
```

## ToSlice

把树上的所有单词和绑定的值以二元组键值对列表的形式返回 

```go
func (x *Trie[T]) ToSlice(delimiter ...string) []*tuple.Tuple2[string, T] 
```

## QueryByPrefix

根据前缀查询单词

```go
func (x *Trie[T]) QueryByPrefix(prefix string, delimiter ...string) []*tuple.Tuple2[string, T]
```

## Contains

查询给定的单词是否在树上 

```go
func (x *Trie[T]) Contains(path string) (bool, error)
```

# TODO 

- 加入导出dot language的支持以方便可视化观察字典树





