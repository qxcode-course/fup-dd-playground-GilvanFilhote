package main
import (
    "fmt"
    "sort"
)
func main() {
    var v int
    fmt.Scan(&v)
    vetor := make([]float64, v)
    for i:=0; i<v; i++{
        fmt.Scan(&vetor[i])

    }
    sort.Float64s(vetor)
    var mediana float64
    if v%2==0{
        mediana=(vetor[v/2-1]+vetor[v/2])/2
    } else {
        mediana=vetor[v/2]
    }

        fmt.Printf("%.1f\n", mediana)

}
    
