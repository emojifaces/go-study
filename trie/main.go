package main

import "fmt"

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// 插入单词到Trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			currentNode.children[ch] = NewTrieNode()
		}
		currentNode = currentNode.children[ch]
	}
	currentNode.isEndOfWord = true // 标记单词结束
}

// 搜索Trie中是否存在单词
func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, ch := range word {
		if node, ok := currentNode.children[ch]; ok {
			currentNode = node
		} else {
			return false
		}
	}
	return currentNode.isEndOfWord // 返回是否为单词的末尾，即是否完全匹配
}

// 检查Trie中是否有以给定前缀开始的任意单词
func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, ch := range prefix {
		if node, ok := currentNode.children[ch]; ok {
			currentNode = node
		} else {
			return false
		}
	}
	return true // 前缀存在
}

func main() {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println(trie.Search("hello"))      // 输出: true
	fmt.Println(trie.Search("world"))      // 输出: true
	fmt.Println(trie.Search("helloo"))     // 输出: false
	fmt.Println(trie.StartsWith("hell"))   // 输出: true
	fmt.Println(trie.StartsWith("worlds")) // 输出: false
}
