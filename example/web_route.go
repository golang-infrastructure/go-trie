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
