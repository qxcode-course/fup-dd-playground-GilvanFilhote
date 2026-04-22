package main
import "fmt"
func main() {
    var p, d1, d2, soma int
    fmt.Scan(&p, &d1, &d2)
    soma=d1+d2
    if p == 0 {
        if soma%2==0{
            fmt.Println(0)
        } else if soma%2!=0{
            fmt.Println(1)
        }
    } else if p == 1 {
        if soma%2==0{
            fmt.Println(1)
        } else if soma%2!=0{
            fmt.Println(0)
        }
    }
}
