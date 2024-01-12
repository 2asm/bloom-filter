package main 

import (
	"fmt"

	"github.com/2asm/bloom-filter/bloomfilter"
)
func main() {
	// error_rate, number of insertions
	bf := bloomfilter.NewBloomFilter(0.001, 100000)
	bf.Add("hi")
	bf.Add("hello there")
	fmt.Printf("%v\n", bf.Contains("hi"))
	fmt.Printf("%v\n", bf.Contains("hello there"))
	fmt.Printf("%v\n", bf.Contains("we"))
}
