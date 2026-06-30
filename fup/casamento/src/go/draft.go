package main
import (
    "fmt"
    "slices"
)
func main() {
    vetor:= make([]int, 5)
    for i:=0; i<5; i++{
        fmt.Scan(&vetor[i])
    }
   slices.Sort(vetor)
    soma:=vetor[0]+vetor[4]
    fmt.Printf("%d\n", soma)
}