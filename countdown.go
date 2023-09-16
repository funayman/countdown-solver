package countdown

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Countdown struct {
	dict *dictionary
}

func (c *Countdown) Solve(chars string) []string {
	rslts := make(chan string, 1)
	words := make([]string, 0)

	go func() {
		c.solver(c.dict.head, []rune(chars), rslts)
		close(rslts)
	}()

	m := make(map[string]struct{})
	for rslt := range rslts {
		if _, ok := m[rslt]; ok {
			continue
		}

		m[rslt] = struct{}{}
		words = append(words, rslt)
	}

	sort.Slice(words, func(i, j int) bool {
		x := words[i]
		y := words[j]
		deltaLength := len(x) - len(y)

		return deltaLength > 0 || (deltaLength == 0 && x < y)
	})

	return words
}

func (c *Countdown) solver(head *node, chars []rune, rslts chan<- string) {
	if head.isWord {
		rslts <- head.path
	}

	if len(chars) == 0 || len(head.edges) == 0 {
		return
	}

	for i, char := range chars {
		if _, ok := head.edges[char]; !ok {
			continue
		}

		newChars := func(i int, runes []rune) []rune {
			dst := make([]rune, len(runes)-1)

			index := 0
			for j, r := range runes {
				if i == j {
					continue
				}
				dst[index] = r
				index++
			}

			return dst
		}(i, chars)

		c.solver(head.edges[char], newChars, rslts)
	}
}

func (c *Countdown) IsWord(word string) bool {
	head := c.dict.head

	for _, r := range word {
		if _, ok := head.edges[r]; !ok {
			return false
		}

		head = head.edges[r]
	}

	return head.isWord
}

// New returns a new instance of Countdown preloaded with all words provided in
// the dictionary file `fn`.
func New(fn string) (*Countdown, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%q): %w", fn, err)
	}
	defer f.Close()

	d := &dictionary{head: &node{edges: make(map[rune]*node)}}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d.AddWord(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return &Countdown{dict: d}, nil
}
