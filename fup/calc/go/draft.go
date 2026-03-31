package main
import "fmt"
func main() {
    var a, b, resultado int
    var op string
    fmt.Scan(&a, &b)
    fmt.Scan(&op)
    if op== "+" {
    resultado = a + b 
    fmt.Printf("%d\n", resultado)
    } else if op=="-" {
        resultado = (a - b)
        fmt.Printf("%d\n", resultado)

    } else if op=="*" {
    resultado = (a * b)
    fmt.Printf("%d\n", resultado)
    } else if op=="/" {
        resultado = (a / b)
        fmt.Printf("%d\n", resultado)
    }
}
