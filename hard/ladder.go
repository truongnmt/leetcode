package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	L := len(beginWord)

	combos := make(map[string][]string)

	starChar := func(s string, i int) string { // charAsteric
		return s[0:i] + "*" + s[i+1:]
	}

	for _, w := range wordList {
		for i := 0; i < L; i++ {
			x := starChar(w, i)         // maskChar
			xs, has := combos[x]
			// fmt.Println(x)
			// fmt.Println(combos)
			if has {
				combos[x] = append(xs, w)
			} else {
				combos[x] = []string{w}
			}
		}

	}

	type pair struct {
		Word  string
		Depth int
		Next  *pair
	}

	type queue struct {
		First *pair
		Last  *pair
	}

	q := new(queue)

	emptyQ := func() bool {
		return q.First == nil
	}

	enqueue := func(p *pair) {
		if emptyQ() {
			q.First = p
			q.Last = p
		} else {
			q.Last.Next = p
			q.Last = p
		}
	}

	dequeue := func() *pair {
		if emptyQ() {
			return nil
		} else {
			x := q.First
			q.First = x.Next
			x.Next = nil
			return x
		}
	}


	enqueue(&pair{Word: beginWord, Depth: 1})

	fmt.Println(combos)

	visited := make(map[string]bool)
	visited[beginWord] = true

	for !emptyQ() {
		p := dequeue()
		word := p.Word
		depth := p.Depth

		
		for i := 0; i < L; i++ {
			k := starChar(word, i)
			edges, has := combos[k]
			fmt.Printf("edges=%s word=%s\n", edges, word)
			if !has {
				continue
			}

			for _, w := range edges {
				if w == endWord {
					return depth + 1
				}

				if _, seen := visited[w]; !seen {
					enqueue(&pair{Word: w, Depth: depth + 1})
					visited[w] = true
				}
			}
		}
	}
	return 0
}
