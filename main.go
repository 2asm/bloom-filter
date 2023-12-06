package main

import (
	"fmt"
	"hash"
	"math"
	"math/rand"

	"github.com/2asm/bloom-filter/bitset"
	"github.com/spaolacci/murmur3"
)

// error
// number of elemnet you are expecting to add
// m bits, k hash functions, n insertions

type BloomFilter struct {
	b      *bitset.BitSet
	hashes []hash.Hash32
}

type IBloomFilter interface {
	Add(s string)
	Contains(s string) bool
}

func NewBloomFilter(error_rate float64, insertions int64) *BloomFilter {
	if error_rate >= 1 || error_rate <= 0 {
		panic("ERROR: error_rate must be between 0 and 1")
	}
	// m := - n*log2(E)/ln(2)
	m := math.Ceil(-float64(insertions) * math.Log2(error_rate) / math.Log(2))

	// k := ln(2)*m/n
	k := math.Ceil(math.Log(2) * m / float64(insertions))

	var hashes []hash.Hash32
	for i := 0; i < int(k); i++ {
		h := murmur3.New32WithSeed(rand.Uint32())
		hashes = append(hashes, h)
	}
	return &BloomFilter{
		b:      bitset.NewBitSet(int(m)),
		hashes: hashes,
	}
}

func (bf *BloomFilter) Add(s string) {
	for _, h := range bf.hashes {
		h.Reset()
		h.Write([]byte(s))
		digest := h.Sum32()
		pos := int(digest % uint32(bf.b.Len()))
		bf.b.Set(pos)
	}
}

func (bf *BloomFilter) Contains(s string) bool {
	for _, h := range bf.hashes {
		h.Reset()
		h.Write([]byte(s))
		digest := h.Sum32()
		pos := int(digest % uint32(bf.b.Len()))
		if bf.b.IsSet(pos) == false {
			return false
		}
	}
	return true
}

func main() {
	// error_rate, number of insertions
	bf := NewBloomFilter(0.01, 1000000)
	bf.Add("hi")
	bf.Add("hello there")
	fmt.Printf("%v\n", bf.Contains("hi"))
	fmt.Printf("%v\n", bf.Contains("hello there"))
	fmt.Printf("%v\n", bf.Contains("we"))
}
