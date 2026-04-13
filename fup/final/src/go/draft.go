package main
import "fmt"
func main() {
    var nota1 float64
    var nota2 float64
    var media1 float64
    var media2 float64
    var notafinal float64

    fmt.Scan(&nota1)
    fmt.Scan(&nota2)
    fmt.Scan(&notafinal)

    media1= (nota1 + nota2)/2
    media2 =(media1 + notafinal)/2

    if media1 >= 7 {
        fmt.Println("aprovado")
    } else if media1 < 4 {
        fmt.Println("reprovado")
    } else if media1 >= 4 && media1 < 7 {
     if media2 >= 5 {
     fmt.Println("aprovado na final")
     } else {
        fmt.Println("reprovado na final")
     }
    } 
}
