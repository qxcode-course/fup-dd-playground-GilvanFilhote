package main
import "fmt"
func main() {
    arr:= []int {9, 8, 6, 4 , 2, 1 }
    fmt.Print("[")
    for posi, valor:= range arr{
        if posi!=0 {
            fmt.Print(", ")
        }
        fmt.Printf ("valor: %d", valor)
    }
     fmt.Print("]\n")
}
