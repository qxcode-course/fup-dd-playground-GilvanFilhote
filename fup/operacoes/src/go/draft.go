package main
import "fmt"
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)
    soma:= num1 + num2
    sub:= num1 - num2
    mult:= num1 * num2
    div:= float32(num1) / float32(num2)
    resto:= float32(num1 % num2)
    fmt.Printf("%d\n", soma)
    fmt.Printf("%d\n", sub)
    fmt.Printf("%d\n", mult)
    fmt.Printf("%.2f\n", div)
    fmt.Printf("%.0f\n", resto)
}
