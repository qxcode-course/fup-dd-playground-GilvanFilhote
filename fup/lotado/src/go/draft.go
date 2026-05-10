package main
import "fmt"
func main() {
    var c, m int
    circula:=0
    fmt.Scan(&c)
    for i:=m; i<100 ; i++{
        fmt.Scan(&m)
        circula+=m
        if circula==0{
            fmt.Println("vazio")
        } else if circula<0 || circula<c{
            fmt.Println("ainda cabe")
        } else if  circula==c || circula<2*c{
            fmt.Println("lotado")
        } else if circula>=2*c{
            fmt.Println("hora de partir")
            break
        }
    }
}
