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

func (this *Trie) getRepresentNode(word string) *TrieNode {
    cur := this.root
    for _, symbol := range word {
        if _, ok := cur.children[symbol]; ok {
            cur = cur.children[symbol]
        } else {
            return nil
        }
    }
    return cur
}

func (this *Trie) Search(word string) bool {
    node := this.getRepresentNode(word)
    return node != nil && node.end
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.getRepresentNode(prefix)
    return node != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */