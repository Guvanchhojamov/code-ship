package trie

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() - Initializes the trie object.
void insert(String word)    - Inserts the string word into the trie.
boolean search(String word) -  Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) -  Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	curr := this
	for i := 0; i < len(word); i++ {
		charIdx := word[i] - 'a'
		if curr.children[charIdx] == nil {
			newNode := Constructor()
			curr.children[charIdx] = &newNode
		}
		curr = curr.children[charIdx]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	curr := this
	for i := 0; i < len(word); i++ {
		charIdx := word[i] - 'a'
		if curr.children[charIdx] == nil {
			return false
		}
		curr = curr.children[charIdx]
	}
	return curr.isEnd == true // must be end of same inserted word.
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	curr := this
	for i := 0; i < len(prefix); i++ {
		charIdx := prefix[i] - 'a'
		if curr.children[charIdx] == nil {
			return false
		}
		curr = curr.children[charIdx]
	}
	return true // no need to be end.
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
