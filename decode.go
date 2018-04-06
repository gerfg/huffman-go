package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func decodeFile(fileName string) {
	fmt.Println("\n --\t  Decode\t  --\n")

	frequency, data := getDecodeData(fileName)

	fmt.Println("\n >Data to string \n")
	dcd := dataToString(data)

	root := huffmanTree(frequency)
	fmt.Println("\t>Decoding String...")
	decodeStringAndCreateFile(dcd, root)
}

func decodeStringAndCreateFile(dcd string, root Node) {
	var bytesToWrite []byte
	var nodeSearch Node = root
	// fmt.Printf("\nRoot: Letter: %d, Freq: %d\n\n", nodeSearch.Letter, nodeSearch.Freq)
	// fmt.Println(dcd)
	for len(dcd) > 2 {
		// fmt.Println(len(dcd))
		// fmt.Printf("é galho?: %v %d\n", nodeSearch.Letter == uint(0), nodeSearch.Letter)
		if nodeSearch.Letter == uint16(257) {
			// fmt.Printf("é zero?: %v\n", dcd[0] == '0')
			if dcd[0] == '0' {
				if nodeSearch.Esq != nil {
					fmt.Printf("0")
					nodeSearch = *nodeSearch.Esq
				}
				dcd = dcd[1:]
			} else {
				// fmt.Printf("é um?: %v\n", dcd[0] == '1')
				if nodeSearch.Dir != nil {
					fmt.Printf("1")
					nodeSearch = *nodeSearch.Dir
				}
				dcd = dcd[1:]
			}
			// fmt.Println(dcd)
		} else {
			fmt.Printf(" - Folha: %v\n", nodeSearch.Letter)
			bytesToWrite = append(bytesToWrite, byte(nodeSearch.Letter))
			nodeSearch = root
			if len(dcd) == 1 {
				// break
			}
		}
	}
	// bytesToWrite = append(bytesToWrite, byte(10))
	ioutil.WriteFile("uncompressed.bin", bytesToWrite, 0644)
}

func getDecodeData(fileName string) (frequency []uint16, data []byte) {
	data, err, size := ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	frequency = make([]uint16, 256)

	dataFrequency := data[:512]
	data = data[512:]
	fmt.Printf("Data: %d - frequency %d  -  SizeTotal: %d\n\n", len(data), len(dataFrequency), size)

	bff := bytes.NewReader(dataFrequency)
	binary.Read(bff, binary.LittleEndian, &frequency)

	// for _, vl := range frequency {
	// 	fmt.Printf("%d - ", vl)
	// }
	return frequency, data
}

func dataToString(data []byte) (decompr string) {
	var sliceCompressor []string = make([]string, len(data))
	// fmt.Printf("\n\n First\n\n")
	// for _, vl := range data {
	// x, _ := fmt.Printf("%0.8b - ", vl)
	// fmt.Printf("  print: %s\n", x)
	// }
	// fmt.Printf("\n\n Convertidos\n\n")
	for idx, vl := range data {
		sliceCompressor[idx] = fmt.Sprintf("%0.8b", vl)
		// fmt.Printf("%s - ", vl)
	}
	fmt.Printf("\n\n Slice\n\n")
	for _, vl := range sliceCompressor {
		// fmt.Printf("%s - ", vl)
		decompr += vl
	}
	// fmt.Printf("\n\n")
	// fmt.Printf("%d - %s\n\n", len(decompr), decompr)

	// for _, vl := range decompr {
	// fmt.Printf("%s - ", string(vl))
	// }
	// fmt.Printf("\n\n")
	return decompr
}
