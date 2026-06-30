package main
import "fmt"
func main() {
    var v int
    fmt.Scan(&v)
    vetor:=make([]int, v)
    for i:=0; i<v; i++{
        fmt.Scan(&vetor[i])
    }
     fmt.Print("[")
    for i:=0; i<v; i++{
        if vetor[i]%2!=0{
            fmt.Printf(" %d", vetor[i])

        }
    }
    fmt.Printf(" ]\n")
    fmt.Print("[")
    for i:=0; i<v; i++{
        if vetor[i]%2==0{
            fmt.Printf(" %d", vetor[i])

        }
    }
    fmt.Printf(" ]\n")
}