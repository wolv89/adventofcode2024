package structures

import "strings"

/*
 * !!!
 * Package is assuming "standard" ASCII characters
 * Ie strings are bytes, not runes
 * String len is equal to number of bytes
 */

type TrieNode struct {
	children map[byte]*TrieNode
	word     bool
}

type Trie struct {
	root    *TrieNode
	longest int
}

func NewTrie() Trie {
	return Trie{
		&TrieNode{
			make(map[byte]*TrieNode),
			false,
		},
		0,
	}
}

func (t Trie) Longest() int {
	return t.longest
}

func (t *Trie) Insert(word string) {

	var (
		c  byte
		ok bool
	)

	node := t.root
	n := len(word)

	t.longest = max(t.longest, n)

	for w := 0; w < n; w++ {
		c = word[w]
		if _, ok = node.children[c]; !ok {
			node.children[c] = &TrieNode{
				make(map[byte]*TrieNode),
				false,
			}
		}
		node = node.children[c]
	}

	node.word = true

}

func (t Trie) Search(word string) bool {

	n := len(word)
	if n > t.longest {
		return false
	}

	var (
		c  byte
		ok bool
	)

	node := t.root

	for w := 0; w < len(word); w++ {
		c = word[w]
		if _, ok = node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return true

}

// Function implemented specifically for day 19 of AOC !
func (t Trie) FindSubmatches(word string) []string {

	var (
		b  strings.Builder
		c  byte
		ok bool
	)

	found := make([]string, 0)

	node := t.root

	for w := 0; w < len(word); w++ {

		c = word[w]

		if _, ok = node.children[c]; !ok {
			break
		}

		b.WriteByte(c)
		node = node.children[c]

		if node.word {
			found = append(found, b.String())
		}

	}

	return found

}
