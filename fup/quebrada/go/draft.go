package main
import "fmt"

func main() {
    var a, b int
    var resto, resultado float32
    fmt.Scan(&a, &b)
    resultado= float32 (a/b)
    resto = float32(a%b)
    fmt.Printf("%.0f\n%.0f", resultado, resto)
    fmt.Printf("%2.f\n", resultado)
}
