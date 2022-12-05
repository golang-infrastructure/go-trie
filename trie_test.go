package trie

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDefaultPathSplitFunc(t *testing.T) {
	slice, err := DefaultPathSplitFunc("中国")
	assert.Nil(t, err)
	t.Log(slice)
}

func TestNewTrie(t *testing.T) {

	// 假装是一个Web路由
	trie := NewTrie[func() error](func(s string) ([]string, error) {
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
	assert.Nil(t, err)

	// 尝试寻找不存在的路由
	value, err := trie.QueryE("/foo")
	assert.NotNil(t, err)
	assert.Nil(t, value)

	// 尝试寻找存在的路由
	handler, err := trie.QueryE("//foo//bar/")
	assert.Nil(t, err)
	assert.NotNil(t, handler)
	err = handler()
	assert.Nil(t, err)

}

func TestNewTrieNode(t *testing.T) {

}

func TestTrieNode_Remove(t *testing.T) {

}

func TestTrieNode_RemoveChild(t *testing.T) {

}

func TestTrie_Add(t *testing.T) {

}

func TestTrie_FindTrieNode(t *testing.T) {

}

func TestTrie_QueryE(t *testing.T) {

}

func TestTrie_QueryOrDefaultE(t *testing.T) {

}

func TestTrie_Remove(t *testing.T) {

}

func TestTrie_Upsert(t *testing.T) {

}
