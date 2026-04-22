package main
import "fmt"
func main() {
    var b, atm int
    var t, atf int 
    fmt.Scan(&b, &t)
    atm = (b + t) * 70/2
    B := 160 - b 
    T := 160 - t
    atf = (B + T ) * 70/2
    if atm > atf {
        fmt.Println(1)
    } else if atf > atm {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
    
}
