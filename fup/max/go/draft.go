package main
import "fmt"
func main() {
    var a, b, maioral float64
    fmt.Scan(&a, &b)
    if a>b{
        maioral = a
    }
    if b>a{
        maioral = b
    }
    if b==a{
        maioral=a
    }
    fmt.Printf("%1.f\n", maioral)
}
