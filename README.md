# Bloom-Filter

Fast bloom filter implementation in golang

### Quick start
``` Golang
func main() {
    // error_rate, number of insertions
    bf := bloomfilter.NewBloomFilter(0.001, 100000)
    bf.Add("dog")
    bf.Add("cat")
    fmt.Printf("%v\n", bf.Contains("dog"))
    fmt.Printf("%v\n", bf.Contains("cat"))
    fmt.Printf("%v\n", bf.Contains("hat"))
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

