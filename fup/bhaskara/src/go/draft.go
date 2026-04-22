package main
import "fmt"
import "math"
func main() {
    var a, b, c, x1, x2 float64
    fmt.Scan(&a, &b, &c)
    delta:= b*b -4*a*c
    if delta>0 {
        x1= (-b + math.Sqrt(delta))/ (2*a)
        x2= (-b - math.Sqrt(delta)) / (2*a)
    fmt.Printf("%.2f\n%.2f\n", x1, x2)
    } else if delta==0 {
        x1= -b/ (2*a) + 0
        fmt.Printf("%.2f\n", x1)
    } else {
        fmt.Printf("nao ha raiz real\n")
    }
    
}
