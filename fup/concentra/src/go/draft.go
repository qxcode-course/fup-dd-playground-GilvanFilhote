package main
import "fmt"
func main() {
    var a, b int 
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i, j:=a, b; i<=b; i, j= i+1, j-1{
        fmt.Printf("%d %d ", i, j)
    }
    fmt.Print("]\n")
}
