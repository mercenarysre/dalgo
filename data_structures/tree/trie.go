package main

// A trie is a tree which is a collection of nodes that point to any number of nodes;
// not a binary tree.
// the root node contains a hash table with the keys "a", "b", and "c". The
// values are other trie nodes, which are the children of this node.

// Trie Struct: keeps track of the root node
// Implementation of a TrieNode: contains a hash table and a word (false/true) attribute which determines
// if a word end at that TrieNode, in this hash table, the keys are individual character strings
// the values are instances/pointers of other TrieNodes

type TrieNode struct {
	Children map[string]*TrieNode
	Word     bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode),
		Word:     false,
	}
}

// When we introduce an instance of a Trie class, we will
// have a new TrieNode as the root node in the constructor
type Trie struct {
	Root *TrieNode
}

// Constructor for Trie
func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

// Trie Insertion
// Time Complexity: O(k)
// we iterate a single time for each character of the target word
// where k is the length of the target word, not data elements of the trie data structure
// Space Complexity:
// We iterate through each character of the word we want to insert,
// if the character does not exist, we can insert it into the trie
// along with its children; Otherwise we can keep tracing down the trie
// once we have reached the last character, we can mark the TrieNode as a word(true)
func (t *Trie) Insert(word string) {
	curr := t.Root
	for _, c := range word {
		ch := string(c) // convert rune to string
		if _, exists := curr.Children[ch]; !exists {
			curr.Children[ch] = NewTrieNode()
		}
		curr = curr.Children[ch]
	}
	curr.Word = true
}

// Trie Search
// Time Complexity: O(k)
// we iterate a single time for each character of the target word
// where k is the length of the target word, not data elements of the trie data structure
// Space Complexity:
// Search whether a word exists in the trie and return a boolean
// we iterate through each character and as soon as we encounter a character
// that is not the trie, we can return false
// It could also be the case that every single character exists but the
// last character is not marked as a word(false)
func (t *Trie) Search(word string) bool {
	curr := t.Root
	for _, c := range word {
		ch := string(c)
		if _, exists := curr.Children[ch]; !exists {
			return false
		}
		curr = curr.Children[ch]
	}
	return curr.Word
}

// Trie SearchPrefix,
// Time Complexity: O(k)
// we iterate a single time for each character of the target word
// where k is the length of the target word, not data elements of the trie data structure
// Space Complexity:
// searching if the trie contains a word prefix("nope", "no"), if each character of
// of the prefix is	within the trie, we return true; Otherwise we return false
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.Root
	for _, c := range prefix {
		ch := string(c)
		if _, exists := curr.Children[ch]; !exists {
			return false
		}
		curr = curr.Children[ch]
	}
	return true
}
