// Package bloomfilter to check item in set
package bloomfilter

import (
	"hash"
	"hash/fnv"
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

// Add item to bloom filter
func (bf *BloomFilter) Add(item []byte) {
	hashvalues := bf.hashValue(item)

	for i := 0; uint(i) < bf.k; i++ {
		pos := uint(hashvalues[i]) % bf.m
		bf.bitset[uint(pos)] = true
	}
	bf.n++
}

// Contain method check if item in the bloom
func (bf *BloomFilter) Contain(item []byte) bool {
	hashvalues := bf.hashValue(item)

	for _, v := range hashvalues {
		pos := uint(v) % bf.m
		if bf.bitset[pos] == false {
			return false
		}
	}
	return true
}

// Use k hashValue func to calc many hashValue values
func (bf *BloomFilter) hashValue(key []byte) (rv []uint64) {
	// TODO: choice better hash functions
	var hashfuncs = []hash.Hash64{fnv.New64(), fnv.New64(), fnv.New64a()}

	for _, hashfunc := range hashfuncs {
		hashfunc.Write(key)
		rv = append(rv, hashfunc.Sum64())
		hashfunc.Reset()
	}
	return
}
