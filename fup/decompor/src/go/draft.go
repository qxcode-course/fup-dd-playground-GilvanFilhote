package main
import (
    "fmt"
    "slices"
)
func main() {
    var n int
    fmt.Scan(&n)
    numeros:= []int{}
    for n>0{
        numeros= append(numeros, n%10)
        n/=10
    }
    slices.Reverse(numeros)
    for i:=0; i<len(numeros); i++{
        fmt.Printf("%d", numeros[i])
        if i!=len(numeros)-1{
            fmt.Printf(" ")
        }
    }
    fmt.Printf("\n")
}