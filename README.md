# Bloom-Filter

Fast bloom filter implementation in golang

### Quick start
``` Golang
func main() {
    // bloom filter will set how many hash functions to use and size of the 
    // bitset based on error_rate and expected number of calls to Add method
	// NewBloomFilter takes error_rate and number of insertions as arguments
    bf := NewBloomFilter(0.01, 1000000)
    bf.Add("hi")
    bf.Add("hello there")
    fmt.Printf("%v\n", bf.Contains("hi"))
    fmt.Printf("%v\n", bf.Contains("hello there"))
    fmt.Printf("%v\n", bf.Contains("we"))
}
```
Output: 
``` Console
true
true
false

[Process exited 0]
```

