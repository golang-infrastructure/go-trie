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
