type TrieNode struct {
    children map[rune]*TrieNode
    end bool
}

func newTrieNode() *TrieNode {
    return &TrieNode{
        children: make(map[rune]*TrieNode),
        end: false,
    }
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{ root: newTrieNode() }
}


func (this *Trie) Insert(word string)  {
    cur := this.root
    for _, symbol := range word {
        if _, ok := cur.children[symbol]; ok {
            cur = cur.children[symbol]
        } else {
            newNode := newTrieNode()
            cur.children[symbol] = newNode
            cur = newNode
        }
    }
    cur.end = true
}


func (this *Trie) Search(word string) bool {
    cur := this.root
    for _, symbol := range word {
        if _, ok := cur.children[symbol]; ok {
            cur = cur.children[symbol]
        } else {
            return false
        }
    }
    return cur.end
}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this.root
    for _, symbol := range prefix {
        if _, ok := cur.children[symbol]; ok {
            cur = cur.children[symbol]
        } else {
            return false
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */