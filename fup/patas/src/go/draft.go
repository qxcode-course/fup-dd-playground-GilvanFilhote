package main
import "fmt"
import "math"
func main() {
    var chico, cebola, patas float64
    var t int
    fmt.Scan(&chico, &cebola, &t)
    vetor:= make([]string, t)
    for i:=0;i<t;i++{
        fmt.Scan(&vetor[i])
        if vetor[i]=="v"{
            patas+=4
        } else if vetor[i]=="c"{
            patas+=4
        } else if vetor[i]=="g"{
            patas+=2
        }
    }
    chicobento:=math.Abs(patas-chico)
    cebolinha:=math.Abs(patas-cebola)
    if chicobento!=cebolinha{
        if patas==cebolinha || cebolinha<chicobento{
            fmt.Printf("%.0f\nCebolinha\n", patas)
        } else if patas==chicobento || cebolinha>chicobento{
            fmt.Printf("%.0f\nChico Bento\n", patas)
        }
    } else {
        fmt.Printf("%.0f\nempate\n", patas)
    }
}