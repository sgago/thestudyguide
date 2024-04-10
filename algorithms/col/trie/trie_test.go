package trie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Add(t *testing.T) {
	trie := New()

	trie.Add("cat")
	trie.Add("cap")
	trie.Add("dog")

	trie.Add("pig", "coo")

	fmt.Println(trie)
}

func TestTrie_Find(t *testing.T) {
	trie := New("cat", "cap", "dog", "dig")

	actual := trie.Find("ca")

	fmt.Println(actual)
}

func TestTrie_Count(t *testing.T) {
	trie := New("cat", "cap", "dog", "dig", "coo")

	actual := trie.Count("ca")

	fmt.Println("len:", actual)
	assert.Equal(t, 3, actual)
}

func TestTrie_CountIgnoresDuplicates(t *testing.T) {
	trie := New("cat", "cat", "cat", "cow")

	actual := trie.Count("cat")

	fmt.Println("len:", actual)
	assert.Equal(t, 1, actual)
}
