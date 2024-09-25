// Package trie (pronounced try) provides a simple implementation of a Trie (prefix tree)
// data structure in Go. Tries are useful for efficiently storing and
// retrieving a collection of strings, especially when dealing with
// operations involving prefixes.
package trie

import (
	"fmt"

	"sgago/thestudyguide/errs"
)

// Trie represents a trie data structure.
type Trie struct {
	children map[string]*Trie // Map to store child nodes based on characters
	count    int              // Count of words with this prefix
}

// New creates a new Trie and adds the provided words.
func New(words ...string) *Trie {
	// Initialize the Trie with an empty root node.
	trie := Trie{
		children: make(map[string]*Trie),
		count:    0,
	}

	// Add the provided words to the Trie.
	trie.Add(words...)

	return &trie
}

// Add inserts words into the Trie.
func (t *Trie) Add(words ...string) *Trie {
	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		if _, err := t.TryFind(word); err == errs.NotFound {
			t.add(word)
		}
	}

	return t
}

func (t *Trie) add(word string) *Trie {
	if len(word) == 0 {
		return t
	}

	// Extract the first character of the word.
	r := string(word[0])

	// If the character is not present in children, create a new Trie node.
	if _, ok := t.children[r]; !ok {
		t.children[r] = New()
	}

	// Increment the count and recursively add the remaining part of the word.
	t.count++
	t.children[r].add(word[1:])

	return t
}

// Find returns the node in the Trie corresponding to the given prefix.
func (t *Trie) Find(prefix string) *Trie {
	curr := t

	for i := 0; i < len(prefix); i++ {
		char := string(prefix[i])

		// Traverse the Trie based on characters in the prefix.
		if next, ok := curr.children[char]; ok {
			curr = next
		} else {
			// Panic if the character in the prefix is not found in the Trie.
			panic(fmt.Sprintln("prefix", prefix, "does not exist"))
		}
	}

	return curr
}

func (t *Trie) TryFind(prefix string) (*Trie, error) {
	curr := t

	for i := 0; i < len(prefix); i++ {
		char := string(prefix[i])

		// Traverse the Trie based on characters in the prefix.
		if next, ok := curr.children[char]; ok {
			curr = next
		} else {
			return nil, errs.NotFound
		}
	}

	return curr, nil
}

// Count returns the number of words with the given prefix in the Trie.
func (t *Trie) Count(prefix string) int {
	if len(prefix) == 0 {
		return t.count
	}

	// Retrieve the node corresponding to the prefix and return its count.
	return t.
		Find(prefix[:len(prefix)-1]).
		count
}

func (t *Trie) TryCount(prefix string) (int, error) {
	if len(prefix) == 0 {
		return t.count, nil
	}

	// Retrieve the node corresponding to the prefix and return its count.
	sub, err := t.TryFind(prefix[:len(prefix)-1])

	if err != nil {
		return -1, err
	}

	return sub.count, nil
}
