package bitset

type BitSet struct {
	len   int
	array []uint64
}

type IBitSet interface {
	IsSet(pos int) bool
	Set(pos int)
	Len() int
	String() string
}

func NewBitSet(len int) *BitSet {
	return &BitSet{
		len:   len,
		array: make([]uint64, (len+63)/64),
	}
}

// pos is 0 base index
func (b *BitSet) Len() int {
	return b.len
}

func (b *BitSet) String() string {
	output := ""
	for _, d := range b.array {
		for i := 0; i < 64; i++ {
            if (d>>i&1) == 1 {
                output += "1"
            } else {
                output += "0"
            }
		}
	}
	return output
}

// pos is 0 base index
func (b *BitSet) IsSet(pos int) bool {
	i := pos / 64
	j := pos % 64
	return (b.array[i] >> j & 1) == 1
}

func (b *BitSet) Set(pos int) {
	i := pos / 64
	j := pos % 64
	b.array[i] |= (1 << j)
}
