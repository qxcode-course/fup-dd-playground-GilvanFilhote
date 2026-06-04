package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    vet_n:=make([]int, n)
    vet_m:=make([]int, m)
    for i:=n; i<n; i++{
        fmt.Scan(&vet_n[i])
    }
     for y:=m; y<m; y++{
        fmt.Scan(&vet_m[y])
    }
     for i:=n; i<n; i++{
        if vet_n[i]==vet_m[i]{
            fmt.Println("sim")
        } else {
            fmt.Println("nao")
        }
    }

    
}