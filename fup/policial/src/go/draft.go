package main
import (
    "fmt"
    "slices"
    ) 
    
func main() {
    var n int
    fmt.Scan(&n)
    vetor:= make([]int, n)
    for i:=0; i<n; i++{
        fmt.Scan(&vetor[i])
    }
    slices.Sort(vetor)
    for i:= 0; i<n; i++{
        fmt.Printf("%d", vetor[i])
        if i!=n-1{
            fmt.Printf(" ")
        }
    }
    fmt.Printf("\n")
}