package main
import "fmt"
func main() {
    var num1, num2, num3, num4 int
    fmt.Scan(&num1, &num2, &num3, &num4)
    if num1>=num2 && num1>num3 && num1>num4 {
        fmt.Printf("%d\n", num1)
    } else if num1<num2 && num2>num3 && num2>num4 {
        fmt.Printf("%d\n", num2)
    } else if num3>=num1 && num2<num3 && num3>num4 {
        fmt.Printf("%d\n", num3)
    } else {
        fmt.Printf("%d\n", num4)
    }
}