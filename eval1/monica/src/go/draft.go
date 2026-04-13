
package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)
    filho:=m-(a+b)
    if a>b && a>filho {
        fmt.Printf("%d\n", a)
    } else if b>a && b>filho {
        fmt.Printf("%d\n", b)
    } else {
        fmt.Printf("%d\n", filho)
}
}
