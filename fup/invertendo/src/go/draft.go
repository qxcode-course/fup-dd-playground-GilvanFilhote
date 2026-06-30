package main
import (
    "fmt"
    "slices"
)
func main() {
    var v int
    fmt.Scan(&v)
    vetor:= make([]int, v)
    for i:=0; i<v; i++{
        fmt.Scan(&vetor[i])
    }

    slices.Reverse(vetor)
    fmt.Printf("[")
    for i:=0; i<v; i++{
        fmt.Printf(" %d", vetor[i])
    }
    fmt.Printf(" ]\n")
}