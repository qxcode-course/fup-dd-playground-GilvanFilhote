package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    vet_n:=make([]int, n)
    vet_m:=make([]int, m)
    for i:=0; i<n; i++{
        fmt.Scan(&vet_n[i])
        for j:=0; j<m; j++{
            fmt.Scan(&vet_m[i])
        }
        fmt.Print("hhh")
    }
}
    