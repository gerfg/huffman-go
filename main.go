package main

import (
	"fmt"
	"sort"
)

func main() {
	frequency, size := getFrequencySlice("t1.bin")

	fmt.Printf("\n%d - Bytes\n", size)

	fmt.Print("\n\n\tInitialize Tree\n\n")
	arrayNodes := initializeNodes(frequency)
	for idx, vl := range arrayNodes {
		fmt.Printf("ID: %d - Letter: %s - Freq: %d - Esq: %v - Dir: %v\n", idx, string(vl.Letter), vl.Freq, vl.Esq, vl.Dir)
	}

	fmt.Print("\n\n\tSort arrayNodes\n\n")

	sort.Slice(arrayNodes, func(i, j int) bool {
		return arrayNodes[i].Freq < arrayNodes[j].Freq
	})
	for idx, vl := range arrayNodes {
		fmt.Printf("ID: %d - Letter: %s - Freq: %d - Esq: %v - Dir: %v\n", idx, string(vl.Letter), vl.Freq, vl.Esq, vl.Dir)
	}

	fmt.Print("\n\n\tHuffman Tree Creation\n\n")
	tree := generateHuffmanTree(arrayNodes)
	fmt.Printf("%d - %s  --  %d -- %d\n", tree.Letter, string(tree.Letter), tree.Letter, tree.Freq)

	fmt.Print("\n\n\tBits Table\n\n")

	// buffer := Codes{bits: "", i: 0, used: false}
	cd := make([]Codes, 256)
	for _, vl := range cd {
		vl.bits += ""
		vl.i = 0
		vl.used = false
	}

	var cds = make(map[uint]string, size)
	generateCodes(tree, cds)

	fmt.Println("\n\n Map: \n")
	for id, vl := range cds {
		fmt.Printf("-> %d %d - %s \n", string(id), id, vl)
	}

}
