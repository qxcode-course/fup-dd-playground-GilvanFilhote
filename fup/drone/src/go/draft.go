package main
import "fmt"
func main() {
    var a, b, c, h, l, area_janela int
    fmt.Scan(&a, &b, &c, &h, &l)
    area_janela=h*l
    if area_janela> a*b{
        fmt.Println("S")
    } else if area_janela> a*c{
        fmt.Println("S")
    } else if area_janela> b*c{
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
