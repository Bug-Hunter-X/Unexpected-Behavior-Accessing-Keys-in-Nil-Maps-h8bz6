```go
package main

import "fmt"

func main() {
    var m map[string]int
    if m == nil {
        fmt.Println("Map is nil") // Handle the nil case appropriately
    } else {
        fmt.Println(m["a"]) 
    }

    // Or initialize the map:
    m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["a"]) // Now it will print 1
}
```