package main
import "fmt"
func main() {
    var a, b, c float32
    fmt.Scan(&a, &b, &c)
    var v1, v2, v3 float32
    fmt.Scan(&v1, &v2, &v3)
    var quantidade float32
    fmt.Scan(&quantidade)
    var dinheiro float32
    dinheiro = (a*v1)+(b*v2)+(c*v3)
    var troco float32
    troco=quantidade-dinheiro
    fmt.Printf("%.2f\n", troco)
}
