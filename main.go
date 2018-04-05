package main

import (
	"fmt"
	"sort"
)

func main() {
	fileName := "t1"
	extension := ".bin"
	encodeFile(fileName + extension)
	decodeFile(fileName + "Encoded" + extension)
}

func encodeFile(fileName string) {
	frequency, size, data := getFrequencySlice(fileName)

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

	createEncodedFile("t1Encoded.bin", compr, frequency)
}

func decodeFile(fileName string) {
	var frequency = make([]int, 256)
	data, err, size := ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(size)

	fmt.Println("\n -- Frequency --\n")
	for i := 0; i < 256; i++ {
		frequency[i] = int(data[i])
		if frequency[i] > 0 {
			fmt.Printf("%d - %d\n", string(i), frequency[i])
		}
	}

	fmt.Println("\n -- Data --\n")
	var dataFile []byte
	for i := 256; i < size; i++ {
		dataFile = append(dataFile[:], data[i])
	}
	fmt.Printf("%0.8b\n", dataFile)

}
