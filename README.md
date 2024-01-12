# Bloom-Filter

Fast bloom filter implementation in golang

### Quick start
``` Golang
func main() {
    // NewBloomFilter takes error_rate and number of insertions as arguments
    bf := bloomfilter.NewBloomFilter(0.001, 100000)
    bf.Add("hi")
    bf.Add("hello there")
    fmt.Printf("%v\n", bf.Contains("hi"))
    fmt.Printf("%v\n", bf.Contains("hello there"))
    fmt.Printf("%v\n", bf.Contains("we"))
}

```
Output: 
``` Console
2024/01/12 22:34:34 Bloom Filter created with 10 Hash Functions and BitSet<1437759>
true
true
false

[Process exited 0]
```

