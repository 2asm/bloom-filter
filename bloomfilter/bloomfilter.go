package bloomfilter

import (
	"log"
	"github.com/2asm/bloom-filter/bitset"
	"github.com/spaolacci/murmur3"
	"math"
)

// error
// number of elemnet you are expecting to add
// m bits, k hash functions, n insertions

type BloomFilter struct {
	FuncCount uint64
	BitCount  uint64
	Bset      *bitset.BitSet
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
    defer log.Printf("Bloom Filter creater where hash function count is %v and bitset size is %d", k, uint64(m))
	return &BloomFilter{
		FuncCount: uint64(k),
		BitCount:  uint64(m),
		Bset:      bitset.NewBitSet(int64(m)),
	}
}

func (bf *BloomFilter) Add(s string) {
	h1, h2 := murmur3.Sum128([]byte(s))
    // shift to fit the hash in 64bit uint
    h1 >>= 8
    h2 >>= 8
    hash := h1
    for i:=uint64(0) ; i<bf.FuncCount; i++ {
        hash += bf.FuncCount*h2
        bf.Bset.Set(int(hash%bf.BitCount))
    }
}

func (bf *BloomFilter) Contains(s string) bool {
	h1, h2 := murmur3.Sum128([]byte(s))
    // shift to fit the hash in 64bit uint
    h1 >>= 8
    h2 >>= 8
    hash := h1
    for i:=uint64(0) ; i<bf.FuncCount; i++ {
        hash += bf.FuncCount*h2
        if !bf.Bset.IsSet(int(hash%bf.BitCount)) {
            return false;
        }
    }
	return true
}
