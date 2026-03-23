package main
import "fmt" 
import "math"

func main() {
    var a float64 
    var b float64
    fmt.Scan(&a, &b)
    var resultado float64
    resultado = math.Abs(a-b)
    fmt.Println(resultado)
}
