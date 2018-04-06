package main

import (
	"fmt"
)

func encodeFile(fileName string) {
	frequency, size, data := getFrequencySlice(fileName)

	fmt.Printf("\n%d - Bytes\n", size)

	fmt.Println("\n --\t  Encode\t  --\n\n")
	tree := huffmanTree(frequency)

	fmt.Print("\n\tBits Table\n")

	var cds = make(map[uint16]string, size)
	generateCodes(tree, cds)

	compr := createEncodeString(data, cds)
	fmt.Println(cds)
	// fmt.Printf("\nlen(compr): %d - %s\n", len(compr), compr)

	createEncodedFile("compressed.bin", compr, frequency)
}
