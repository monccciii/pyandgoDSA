package main

import (
	"fmt"
)

// TrieNode represents a single node in a Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// Trie represents a Trie data structure
type Trie struct {
	root *TrieNode
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, r := range word {
		if current.children[r] == nil {
			current.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[r]
	}
	current.isWord = true
}

// Autocomplete returns all words in the Trie that start with prefix
func (t *Trie) Autocomplete(prefix string) []string {
	current := t.root
	for _, r := range prefix {
		if current.children[r] == nil {
			return nil
		}
		current = current.children[r]
	}
	return t.findWords(current, prefix)
}

func (t *Trie) findWords(node *TrieNode, prefix string) []string {
	var words []string
	if node.isWord {
		words = append(words, prefix)
	}
	for r, child := range node.children {
		words = append(words, t.findWords(child, prefix+string(r))...)
	}
	return words
}

func main() {
	trie := &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
	trie.Insert("dog")
	trie.Insert("deer")
	trie.Insert("deal")
	trie.Insert("dear")
	trie.Insert("cat")
	trie.Insert("car")
	trie.Insert("care")
	trie.Insert("cow")
	trie.Insert("camel")
	trie.Insert("horse")
	fmt.Println(trie.Autocomplete("ca")) // Output: [car care]
}
