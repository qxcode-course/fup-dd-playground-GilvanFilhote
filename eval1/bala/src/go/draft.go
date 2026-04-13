package main
import "fmt"
import "math"
func main() {
    var ponto1x, ponto1y, ponto2x, ponto2y float64
    fmt.Scan(&ponto1x, &ponto1y, &ponto2x, &ponto2y)
    distancia:= math.Sqrt{(ponto2x - ponto1x)^2 + (ponto2y-ponto1y)^2}
    fmt.Printf("%.2f", distancia)
}
