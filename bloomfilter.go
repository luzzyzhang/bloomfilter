// Package bloomfilter to check item in set
package bloomfilter

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
)

// BloomFilter data strucutre define
type BloomFilter struct {
	k      uint   // Number of hash values
	n      uint   // Number of elements in the filter
	m      uint   // Size of the bloom filter bitset
	bitset []bool // The bloom-filter bitset
}

// New make a new BloomFilter struct
func New(size uint) *BloomFilter {
	return &BloomFilter{
		k:      3,
		m:      size,
		n:      uint(0),
		bitset: make([]bool, size),
	}
}

var hashfunc = sha1.New()

func (bf *BloomFilter) add(s string) {
	index := bf.position(s)
	bf.bitset[index] = true
}

func (bf *BloomFilter) query(s string) bool {
	index := bf.position(s)
	// fmt.Printf("s: %v, index: %v\n", s, index)
	fmt.Printf("bf.bits: %v", bf.bitset)
	return bf.bitset[index]
}

func (bf *BloomFilter) position(s string) int {
	hs := hasher(hashfunc, s)
	if hs < 0 {
		hs = -hs
	}
	return hs % len(bf.bitset)
}

func hasher(h hash.Hash, s string) int {
	bits := h.Sum([]byte(s))
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result)
}
