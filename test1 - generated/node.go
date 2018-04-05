package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type Node struct {
	Letter uint
	Freq   int
	Esq    *Node
	Dir    *Node
}

type Code struct {
	Code byte
}

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
	var n Node
	for len(arrayNodes) > 1 {
		fmt.Println(len(arrayNodes))
		for _, vl := range arrayNodes {
			fmt.Printf("%v\n", vl)
		}
		// arrayNodes[0] + arrayNodes[1]
		n = createNode(&arrayNodes[0], &arrayNodes[1])
		arrayNodes[1] = n
		arrayNodes = append(arrayNodes[:0], arrayNodes[1:]...)
		sort.Slice(arrayNodes, func(i, j int) bool {
			if arrayNodes[i].Freq == arrayNodes[j].Freq {
				return arrayNodes[i].Letter < arrayNodes[j].Letter
			} else {
				return arrayNodes[i].Freq < arrayNodes[j].Freq
			}
		})
	}
	return arrayNodes[0]
}

func createNode(n1 *Node, n2 *Node) (n Node) {
	n = Node{Letter: 0, Freq: n1.Freq + n2.Freq, Esq: &Node{Letter: n1.Letter, Freq: n1.Freq, Esq: n1.Esq, Dir: n1.Dir}, Dir: &Node{Letter: n2.Letter, Freq: n2.Freq, Esq: n2.Esq, Dir: n2.Dir}}
	return n
}

func createEncodeString(data []byte, codes map[uint]string) (compressed string) {
	for _, vl := range data {
		compressed += codes[uint(vl)]
		// fmt.Printf("%v - %v -> %v - %v\n", vl, string(vl), codes[uint(vl)], compressed)
	}
	return compressed
}

func createEncodedFile(fileName string, compress string, frequency []int) {
	var bt2 uint8
	var bitsBuffer = 0

	numberOfBytes := math.Floor(float64(len(compress)) / 8)
	bytesCreated := 0
	lastBits := len(compress) % 8
	fmt.Printf("\n-- Byte's: %d  --  lastBits: %d\n\n", numberOfBytes, lastBits)

	var bytesToWrite []byte

	for _, vl := range compress {
		if vl == '0' {
			// fmt.Printf("%d Bit 0: %b\n", idx, vl)
			bt2 = bt2 << 1
		}
		if vl == '1' {
			// fmt.Printf("%d Bit 1: %b\n", idx, vl)
			bt2 = bt2<<1 + 1
		}
		bitsBuffer++
		if bitsBuffer == 8 {
			bytesCreated++
			// fmt.Printf("\nBuffer 8bits: %s %0.8b\n\n", bt2, bt2)
			bytesToWrite = append(bytesToWrite[:], bt2)
			bitsBuffer = 0
			bt2 = 0
		}
	}
	for i := 0; i < (8 - lastBits); i++ {
		bt2 = bt2 << 1
		// fmt.Printf("%d Bit 0:\n", i)
	}
	bytesCreated++
	bytesToWrite = append(bytesToWrite[:], bt2)
	// fmt.Printf("\nBuffer 8bits: %s %0.8b\n\n", bytesCreated, bt2, bt2)

	// fmt.Printf("\n\n-- Array of bytes --\n\n")
	// for idx, vl := range bytesToWrite {
	// 	fmt.Printf("%d  -  %0.8b  -  %s\n", idx, vl, vl)
	// }

	err := ioutil.WriteFile(fileName, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

}

func generateCodes(tree Node, cds map[uint]string) {

	var walkTree func(n *Node, code string, cds map[uint]string)

	walkTree = func(n *Node, code string, cds map[uint]string) {
		if n.Esq == nil {
			fmt.Printf("'%s' - %d -> %s\n", string(n.Letter), n.Freq, code)
			cds[n.Letter] = code
			return
		}
		code += "0"
		walkTree(n.Esq, code, cds)
		code = code[:len(code)-1]
		code += "1"
		walkTree(n.Dir, code, cds)
		code = code[:len(code)-1]
	}
	var code string
	walkTree(&tree, code, cds)
}

func showPreOrder(tree *Node) {
	fmt.Printf("root: %v \n", tree)

	var preOrder func(tree *Node)
	preOrder = func(tree *Node) {

		if (*tree).Esq == nil {
			fmt.Printf("%v \n", *tree)
			return
		}

		if (*tree).Esq != nil {
			preOrder(tree.Esq)
		}
		if (*tree).Dir != nil {
			preOrder(tree.Dir)
		}

	}
	preOrder(tree.Esq)
	preOrder(tree.Dir)
}
