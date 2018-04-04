package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile(fileName string) (data []byte, err error, size int) {
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	return data, err, len(data)
}

func getFrequencySlice(fileName string) (frequency []int, size int) {
	frequency = make([]int, 256)
	for i := range frequency {
		frequency[i] = 0
	}

	data, _, size := ReadFile(fileName)
	for idx, vl := range data {
		fmt.Printf("Value: %s  -  %s  --  index: %d\n", vl, string(vl), idx)
		frequency[vl]++
	}

	fmt.Println("\n\n\t Frequencia")
	for idx, vl := range frequency {
		if vl > 0 {
			fmt.Printf("Char: %s  ->  %d\n", string(idx), vl)
		}
	}
	return frequency, size
}
