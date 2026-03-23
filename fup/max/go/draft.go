package main
import ("fmt"
"math"
)
func main() {
    var a, b, maioral float64
    fmt.Scan(&a, &b)
    maioral = math.Max(a, b)
    fmt.Println(maioral)
    
}
