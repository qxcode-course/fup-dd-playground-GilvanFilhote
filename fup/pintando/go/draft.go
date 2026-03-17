package main
import ("fmt" 
        "math"
)
func main() {
    var at, bt, ct, a, p, resultado float64
    fmt.Scan(&at, &bt, &ct)
    p = (at + bt + ct) / 2
    a = p*(p-at)*(p-bt)*(p-ct)
    resultado = math.Sqrt(a)
    fmt.Printf("%.2f\n", resultado)

}
