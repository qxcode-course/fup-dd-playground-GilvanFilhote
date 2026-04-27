package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Printf("[ ")
    for i := 0; i <= 9; i++ {
        if i == n { continue }
        fmt.Printf("%d ", i)
    }
    if n!= 10{
    fmt.Println("ceu ]")
    } else {
    fmt.Println("]")
}
}
