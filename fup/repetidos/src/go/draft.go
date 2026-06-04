package main
import "fmt"
func main() {
    var p, n, rep int
    fmt.Scan(&p, &n)
    vet:=make([]int, n)
    for i:=0; i<n; i++{
        fmt.Scan(&vet[i])
        if vet[i]==p{
            rep++
        }
    }
    fmt.Println(rep)
}