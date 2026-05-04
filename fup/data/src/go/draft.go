package main
import "fmt"
func main() {
    var hora, minuto, dia, mês, ano int
    fmt.Scan(&hora, &minuto, &dia, &mês, &ano)
    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", hora, minuto, dia, mês, ano%100)
}
