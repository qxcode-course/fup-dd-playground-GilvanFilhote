package main
import "fmt"
func main() {
    var a, b, soma_2e3 int
    fmt.Scan(&a, &b)
    if a<b{
        for i:=a; i<b; i++ {
            if i%2==0 && i%3==0{
            soma_2e3 += i
            }
        }
        fmt.Printf("%d\n", soma_2e3)
    } else if a==b{
         fmt.Printf("%d\n", a)
    } else {
         fmt.Printf("invalido\n")
    }
}
