package main
import "fmt"
import "math"
func main() {
    var operacao string
    var nota float64
    fmt.Scan(&operacao, &nota)
    if operacao == "r" {
        fmt.Println(int(math.Round(nota)))
        } else if operacao == "c"  {
            fmt.Println(int(math.Ceil(nota)))
        } else{ 
            fmt.Println(int(math.Floor(nota))) // para floor
}
}