package main
import "fmt"
func main() {
    var hora, minuto, dia, mês, ano int
    fmt.Scan(&hora, &minuto, &dia, &mês, &ano)
    fmt.Printf("0%d:%d %d/0%d/%d\n", hora, minuto, dia, mês, ano)
}
