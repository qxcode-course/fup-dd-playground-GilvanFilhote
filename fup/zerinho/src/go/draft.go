package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    if a==b && a==c{
        fmt.Println("empate")
    } else if a==b && b!=c{
         fmt.Println("jog3")
    } else if a!=b && b==c{
         fmt.Println("jog1")
    } else {
        fmt.Println("jog2")
    }
}