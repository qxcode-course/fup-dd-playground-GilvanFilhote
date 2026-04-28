package main
import "fmt"
func main() {
    var in, f int
    fmt.Scan(&in, &f)
    for i := in; i <= f ; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("zigzag")
        } else if i % 3 == 0 {
            fmt.Println("zig")
        } else if i % 5 == 0 {
            fmt.Println("zag")
        } else if i % 3 != 0 && i % 5 != 0 {
            fmt.Println(i)
        }
    }
}

