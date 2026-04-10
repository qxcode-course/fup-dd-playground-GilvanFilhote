package main
import "fmt"
func main() {
    var a, b, soma_par int 
    fmt.Scan(&a, &b)
    if a<b{
        for i:=a; i<=b; i++{
        if i%2==0{
            soma_par+=i
        }    
        }
        fmt.Printf("%d\n", soma_par)
    } else if a==b {
        fmt.Printf("%d\n", a)
    } else {
        fmt.Printf("invalido\n")
    }
}
