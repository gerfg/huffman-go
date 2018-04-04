package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	Letter uint
	Freq   int
	Esq    *Node
	Dir    *Node
}

// var codes map[string]string

func (n Node) initializeNode(l uint, f int, e *Node, d *Node) {
	n = Node{l, f, e, d}
}

func initializeNodes(frequency []int) (arrayNodes []Node) {
	for i := 0; i < 256; i++ {
		if frequency[i] > 0 {
			arrayNodes = append(arrayNodes, Node{Letter: uint(i), Freq: frequency[i], Esq: nil, Dir: nil})
		}
	}
	return arrayNodes
}

func generateHuffmanTree(arrayNodes []Node) (tree Node) {
	for len(arrayNodes) > 1 {
		fmt.Println(len(arrayNodes))
		for idx, vl := range arrayNodes {
			fmt.Printf("ID: %d - Letter: %s - Freq: %d - Esq: %v - Dir: %v\n", idx, string(vl.Letter), vl.Freq, vl.Esq, vl.Dir)
		}
		// arrayNodes[0] + arrayNodes[1]
		arrayNodes[1] = insertHuffmanTree(arrayNodes[0], arrayNodes[1])
		arrayNodes = append(arrayNodes[:0], arrayNodes[1:]...)
		sort.Slice(arrayNodes, func(i, j int) bool {
			return arrayNodes[i].Freq < arrayNodes[j].Freq
		})
	}
	return arrayNodes[0]
}

func insertHuffmanTree(n1 Node, n2 Node) (n Node) {
	n = Node{Letter: 257, Freq: n1.Freq + n2.Freq, Esq: &n1, Dir: &n2}
	return n
}

func generateCodes(tree Node, cds map[uint]string) {

	var traverse func(n *Node, code uint64, bits byte, cds map[uint]string)

	traverse = func(n *Node, code uint64, bits byte, cds map[uint]string) {
		if n.Esq == nil || n.Dir == nil {
			// Leaf
			fmt.Printf("'%c': %0"+strconv.Itoa(int(bits))+"b  freq: %d\n", n.Letter, code, n.Freq)
			cds[n.Letter] = strconv.Itoa(int(bits))
			return
		}
		bits++
		traverse(n.Esq, code<<1, bits, cds)
		traverse(n.Dir, code<<1+1, bits, cds)
	}
	traverse(&tree, 0, 0, cds)
}
