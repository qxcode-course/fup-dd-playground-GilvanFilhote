package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i:=a; i<b; i++{
        if i%2==0{
            continue
        }
        fmt.Printf("%d ", i)
    }
    fmt.Print("]\n")
}
