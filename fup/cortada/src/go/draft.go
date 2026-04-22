package main
import "fmt"
func main() {
    var B, T int
    fmt.Scan(&B, &T)
    area_trapezio_felix:=(B+T)*70/2
    b:=160-B
    t:=160-T
    area_trapezio_marzia:=(b+t)*70/2
    if area_trapezio_felix>area_trapezio_marzia {
        fmt.Println(1)
    } else if area_trapezio_felix<area_trapezio_marzia {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}
