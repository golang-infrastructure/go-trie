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
	trie := New[func() error](func(s string) ([]string, error) {
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
	value, err := trie.Query("/foo")
	assert.NotNil(t, err)
	assert.Nil(t, value)

	// 尝试寻找存在的路由
	handler, err := trie.Query("//foo//bar/")
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
	trie := New[string]()

	err := trie.Add("test", "测试")
	assert.Nil(t, err)

	slice := trie.ToSlice()
	assert.Equal(t, 1, len(slice))
}

func TestTrie_FindTrieNode(t *testing.T) {

}

func TestTrie_Query(t *testing.T) {
	trie := New[string]()

	err := trie.Add("test", "测试")
	assert.Nil(t, err)

	value, err := trie.Query("test")
	assert.Nil(t, err)
	assert.Equal(t, "测试", value)

	value, err = trie.Query("test-001")
	assert.ErrorIs(t, ErrNotFound, err)
}

func TestTrie_QueryOrDefault(t *testing.T) {
	trie := New[string]()

	err := trie.Add("test", "测试")
	assert.Nil(t, err)

	value, err := trie.QueryOrDefault("test", "策士")
	assert.Nil(t, err)
	assert.Equal(t, "测试", value)

	value, err = trie.QueryOrDefault("test-001", "策士")
	assert.Nil(t, err)
	assert.Equal(t, "策士", value)

}

func TestTrie_Remove(t *testing.T) {
	trie := New[string]()

	err := trie.Upsert("china", "瓷器")
	assert.Nil(t, err)
	err = trie.Upsert("chinese", "中国人")
	assert.Nil(t, err)
	slice := trie.ToSlice()
	assert.Equal(t, 2, len(slice))

	err = trie.Remove("china")
	assert.Nil(t, err)
	slice = trie.ToSlice()
	assert.Equal(t, 1, len(slice))
}

func TestTrie_Upsert(t *testing.T) {
	trie := New[string]()

	_ = trie.Upsert("china", "瓷器")
	value, err := trie.Query("china")
	assert.Nil(t, err)
	assert.Equal(t, "瓷器", value)

	_ = trie.Upsert("china", "中国")
	value, err = trie.Query("china")
	assert.Nil(t, err)
	assert.Equal(t, "中国", value)
}

func TestTrie_ToSlice(t *testing.T) {
	trie := New[string]()
	err := trie.Add("china", "中国")
	assert.Nil(t, err)
	err = trie.Add("chinese", "中国人")
	assert.Nil(t, err)
	slice := trie.ToSlice("")
	assert.Equal(t, 2, len(slice))
}

func TestTrie_FindByPrefix(t *testing.T) {
	trie := New[string]()
	_ = trie.Upsert("china", "china")
	_ = trie.Upsert("chinese", "chinese")
	_ = trie.Upsert("channel", "channel")
	_ = trie.Upsert("chan", "chan")
	_ = trie.Upsert("boy", "boy")
	_ = trie.Upsert("cc11001100", "CC11001100")

	slice := trie.QueryByPrefix("chan")
	//t.Log(slice)
	assert.Equal(t, 2, len(slice))

}

func TestTrie_Contains(t *testing.T) {
	trie := New[string]()
	_ = trie.Upsert("china", "china")
	_ = trie.Upsert("chinese", "chinese")
	_ = trie.Upsert("channel", "channel")
	_ = trie.Upsert("chan", "chan")
	_ = trie.Upsert("boy", "boy")
	_ = trie.Upsert("cc11001100", "CC11001100")

	//t.Log(trie.ExportToDotLanguage())
	// digraph G1 {
	//            "0::" -> "1:c:";
	//            "0::" -> "2:b:";
	//            "1:c:" -> "3:h:";
	//            "1:c:" -> "4:c:";
	//            "2:b:" -> "5:o:";
	//            "3:h:" -> "6:i:";
	//            "3:h:" -> "7:a:";
	//            "4:c:" -> "8:1:";
	//            "5:o:" -> "9:y:boy";
	//            "6:i:" -> "10:n:";
	//            "7:a:" -> "11:n:chan";
	//            "8:1:" -> "12:1:";
	//            "10:n:" -> "13:a:china";
	//            "10:n:" -> "14:e:";
	//            "11:n:chan" -> "15:n:";
	//            "12:1:" -> "16:0:";
	//            "14:e:" -> "17:s:";
	//            "15:n:" -> "18:e:";
	//            "16:0:" -> "19:0:";
	//            "17:s:" -> "20:e:chinese";
	//            "18:e:" -> "21:l:channel";
	//            "19:0:" -> "22:1:";
	//            "22:1:" -> "23:1:";
	//            "23:1:" -> "24:0:";
	//            "24:0:" -> "25:0:CC11001100";
	//        }
}
