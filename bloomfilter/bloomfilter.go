package bloomfilter

import (
	"log"
	"math"
	"math/rand"

	"github.com/2asm/bloom-filter/bitset"
	"github.com/spaolacci/murmur3"
)

// error
// number of elemnet you are expecting to add
// m bits, k hash functions, n insertions

type BloomFilter struct {
	FuncCount uint64
	BitCount  uint64
	Bset      *bitset.BitSet
	Seeds     []uint32
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

	seeds := make([]uint32, int(k))
	for i := 0; i < int(k); i++ {
		seeds[i] = rand.Uint32()
	}
	defer log.Printf("Bloom Filter created with %v Hash Functions and BitSet<%v>", k, uint64(m))
	return &BloomFilter{
		FuncCount: uint64(k),
		BitCount:  uint64(m),
		Bset:      bitset.NewBitSet(int64(m)),
		Seeds:     seeds,
	}
}

func (bf *BloomFilter) Add(s string) {
	for i := uint64(0); i < bf.FuncCount; i++ {
		hash := murmur3.Sum64WithSeed([]byte(s), bf.Seeds[i])
		bf.Bset.Set(int(hash % bf.BitCount))
	}
}

func (bf *BloomFilter) Contains(s string) bool {
	for i := uint64(0); i < bf.FuncCount; i++ {
		hash := murmur3.Sum64WithSeed([]byte(s), bf.Seeds[i])
		if !bf.Bset.IsSet(int(hash % bf.BitCount)) {
			return false
		}
	}
	return true
}
