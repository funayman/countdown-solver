package countdown

import (
	"fmt"
)

type node struct {
	ch     rune
	edges  map[rune]*node
	isWord bool
	path   string
}

type dictionary struct {
	head *node
	size int
}

func (d *dictionary) AddWord(word string) {
	head := d.head

	for _, r := range word {
		if n, ok := head.edges[r]; ok {
			head = n
			continue
		}

		path := fmt.Sprintf("%s%c", head.path, r)
		n := &node{
			ch:     r,
			edges:  make(map[rune]*node),
			isWord: (path == word),
			path:   path,
		}

		head.edges[r] = n
		head = n
	}

	d.size++
}

func (d *dictionary) Count() int {
	return d.size
}
