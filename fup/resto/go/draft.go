package main
import "fmt"
func main() {
    var n, d int
    fmt.Scan(&n, &d)
    var quociente, resto int
    quociente = n/d
    resto = n%d
    fmt.Printf("%d %d\n", quociente, resto)
}
