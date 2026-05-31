package main
import "fmt"
func main() {
    var tam int
    var soma float64
    fmt.Scan(&tam)

    vetor:=make([]float64, tam)
    for i:=0; i<tam; i++{
        fmt.Scan(&vetor[i])
        soma+=vetor[i]
    }
    soma=soma/float64(tam)
    fmt.Printf("%.1f\n", soma)
}