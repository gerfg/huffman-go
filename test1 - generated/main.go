package main

import (
	"fmt"
	"sort"
)

func main() {
	frequency, size, data := getFrequencySlice("generated.bin")

	for _, vl := range data {
		fmt.Println(vl)
	}

	fmt.Printf("\n%d - Bytes\n", size)

	fmt.Print("\n\n\tInitialize Tree\n\n")
	arrayNodes := initializeNodes(frequency)
	// for idx, vl := range arrayNodes {
	// 	fmt.Printf("ID: %d - Letter: %s - Freq: %d - Esq: %v - Dir: %v\n", idx, string(vl.Letter), vl.Freq, vl.Esq, vl.Dir)
	// }

	fmt.Print("\n\n\tSort arrayNodes\n\n")

	sort.Slice(arrayNodes, func(i, j int) bool {
		if arrayNodes[i].Freq == arrayNodes[j].Freq {
			return arrayNodes[i].Letter < arrayNodes[j].Letter
		} else {
			return arrayNodes[i].Freq < arrayNodes[j].Freq
		}
	})
	// for idx, vl := range arrayNodes {
	// 	fmt.Printf("ID: %d - Letter: %s - Freq: %d - Esq: %v - Dir: %v\n", idx, string(vl.Letter), vl.Freq, vl.Esq, vl.Dir)
	// }

	fmt.Print("\n\n\tHuffman Tree Creation\n\n")
	tree := generateHuffmanTree(arrayNodes)
	fmt.Printf("\n\n%d - %s  --  %d  ------ %v - %v\n\n", tree.Letter, string(tree.Letter), tree.Freq, tree.Esq, tree.Dir)
	// fmt.Printf("\n\n%v\n", (*tree.Esq))

	fmt.Print("\n\n PreOrder \n\n")
	showPreOrder(&tree)

	fmt.Print("\n\n\tBits Table\n\n")

	var cds = make(map[uint]string, size)
	generateCodes(tree, cds)

	fmt.Println("\n\n Map: \n")
	// for id, vl := range cds {
	// 	fmt.Printf("-> %s - %s - %s \n", id, string(id), vl)
	// }

	compr := createEncodeString(data, cds)
	fmt.Printf("\n%d - %s\n", len(compr), compr)

	createEncodedFile("t1Enconded.bin", compr, frequency)
}
