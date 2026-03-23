package main
import "fmt"
import "math"
func main() {
    var x1, x2, y1, y2 float64
    fmt.Scan(&x1, &y1, &x2, &y2)
    var distancia_x, distancia_y, distancia_real float64
    distancia_x = (x2-x1)
    distancia_y = (y2-y1)
    distancia_x= math.Pow(distancia_x, 2)
    distancia_y= math.Pow(distancia_y, 2)
    distancia_real= distancia_y+distancia_x
    distancia_real= math.Sqrt(distancia_real)
    fmt.Printf("%.2f\n", distancia_real)
}
