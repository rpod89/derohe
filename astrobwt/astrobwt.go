package astrobwt

import "fmt"
import "golang.org/x/crypto/sha3"
import "golang.org/x/crypto/salsa20/salsa"

// see here to improve the algorithms more https://github.com/y-256/libdivsufsort/blob/wiki/SACA_Benchmarks.md

var x = fmt.Sprintf

const stage1_length int = 9973 // it is a prime

func POW16(inputdata []byte) (outputhash [32]byte) {

	var output [stage1_length]byte
	var counter [16]byte

	key := sha3.Sum256(inputdata)

	var stage1 [stage1_length]byte // stages are taken from it
	salsa.XORKeyStream(stage1[:stage1_length], stage1[:stage1_length], &counter, &key)

	var sa [stage1_length]int16
	text_16_0alloc(stage1[:], sa[:])

	for i := range sa {
		output[i] = stage1[sa[i]]
	}

	//	fmt.Printf("input %+v\n",inputdata)
	//	fmt.Printf("sa %+v\n",sa)
	outputhash = sha3.Sum256(output[:])

	return
}

func text_16_0alloc(text []byte, sa []int16) {
	if int(int16(len(text))) != len(text) || len(text) != len(sa) {
		panic("suffixarray: misuse of text_16")
	}
	var memory [2 * 256]int16
	sais_8_16(text, 256, sa, memory[:])
}

func POW32(inputdata []byte) (outputhash [32]byte) {
	var output [stage1_length]byte
	var counter [16]byte
	key := sha3.Sum256(inputdata)

	var stage1 [stage1_length]byte // stages are taken from it
	salsa.XORKeyStream(stage1[:stage1_length], stage1[:stage1_length], &counter, &key)
	var sa [stage1_length]int32
	text_32_0alloc(stage1[:], sa[:])
	for i := range sa {
		output[i] = stage1[sa[i]]
	}
	outputhash = sha3.Sum256(output[:])

	return
}

func text_32_0alloc(text []byte, sa []int32) {
	if int(int16(len(text))) != len(text) || len(text) != len(sa) {
		panic("suffixarray: misuse of text_16")
	}
	var memory [2 * 256]int32
	sais_8_32(text, 256, sa, memory[:])
}
