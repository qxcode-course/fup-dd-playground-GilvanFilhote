package main
import "fmt"

func main() {
    var a, b int
    var resto, resultado float32
    fmt.Scan(&a, &b)
    resultado = float32 (a) / float32 (b)
    resto = float32(a%b)
    fmt.Printf("%.0f\n", resultado)
    fmt.Printf("%.0f\n", resto)
    fmt.Printf("%.2f\n", resultado)
}
