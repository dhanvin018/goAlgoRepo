package dataStructures

import (
	"fmt"
	"strings"
	"unicode"
)

type Trie struct {
	value      string
	children   []*Trie
	childCount int
}

func NewTrie(char string) *Trie {
	return &Trie{
		value:      char,
		children:   make([]*Trie, 0, 27),
		childCount: 0,
	}
}

func CreateTrie() *Trie {
	return &Trie{
		value:      "*",
		children:   make([]*Trie, 0, 27),
		childCount: 0,
	}
}

func (trie *Trie) Insert(word string) (*Trie, error) {
	if len(word) == 0 {
		return trie, nil
	}
	word = strings.ToLower(word)
	current := trie
	for _, char := range word {
		if !unicode.IsLetter(char) {
			return nil, fmt.Errorf("invalid letter in word")
		}
		index := char - 'a'

		if current.children[index] != nil {
			current = current.children[index]
			continue
		}
		current.childCount += 1
		child := NewTrie(string(char))
		current.children[index] = child
		current = child
	}
	current.childCount += 1
	current.children[26] = NewTrie("/")
	return trie, nil
}

func (trie *Trie) Search(word string) (bool, error) {
	if len(word) == 0 {
		return false, nil
	}
	word = strings.ToLower(word)
	current := trie
	for _, char := range word {
		if !unicode.IsLetter(char) {
			return false, fmt.Errorf("invalid letter in word")
		}
		index := char - 'a'
		if current.children[index] == nil {
			return false, nil
		}
		current = current.children[index]
	}
	if current.children[26] != nil {
		return true, nil
	}
	return false, nil
}

func (trie *Trie) PrefixSearch(word string) (bool, error) {
	if len(word) == 0 {
		return true, nil
	}
	word = strings.ToLower(word)
	current := trie
	for _, char := range word {
		if !unicode.IsLetter(char) {
			return false, fmt.Errorf("invalid letter in word")
		}
		index := char - 'a'
		if current.children[index] == nil {
			return false, nil
		}
		current = current.children[index]
	}
	return true, nil
}

func (trie *Trie) DeleteWord(word string) (*Trie, error) {
	if len(word) == 0 {
		if trie.children[26] != nil {
			trie.children[26] = nil
			trie.childCount -= 1
		}
		return trie, nil
	}
	word = strings.ToLower(word)
	if word[0] < 'a' || word[0] > 'z' {
		return nil, fmt.Errorf("invalid letter in word")
	}
	childIndex := word[0] - 'a'
	if trie.children[childIndex] == nil {
		return nil, fmt.Errorf("word does not exist in trie")
	}
	childRes, error := trie.children[childIndex].DeleteWord(word[1:])
	if error != nil {
		return nil, error
	}
	if childRes.childCount == 0 {
		trie.children[childIndex] = nil
		trie.childCount -= 1
	} else {
		trie.children[childIndex] = childRes
	}
	return trie, nil
}
