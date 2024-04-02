package main 

import (
	"fmt"

	"github.com/2asm/bloom-filter/bloomfilter"
)
func main() {
	// error_rate, number of insertions
	bf := bloomfilter.NewBloomFilter(0.001, 100000)
	bf.Add("dog")
	bf.Add("cat")
	fmt.Printf("%v\n", bf.Contains("dog"))
	fmt.Printf("%v\n", bf.Contains("cat"))
	fmt.Printf("%v\n", bf.Contains("hat"))
}
