## BiMap

A simple bi-directional map implementation in Go

API docs: https://godoc.org/github.com/diogobeda/bimap

## Usage

```go
import "githib.com/diogobeda/bimap"

// empty map initialization
bm := NewBiMap()

bm.Set("a", 1)

fmt.Println(bm.Get("a"))
// 1

fmt.Println(bm.GetKey(1))
// "a"


// Initialize with values
bm := NewBiMap(Tuple{"a", 1}, Tuple{"b", 2})
```