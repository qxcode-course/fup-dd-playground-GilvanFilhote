package main
import "fmt"
func main() {
    var valor_da_tv, parcelas float64
    fmt.Scan(&valor_da_tv, &parcelas)

    if parcelas>1{
        juros:=(parcelas-1)*0.05
        valor_da_tv=valor_da_tv*(1+juros)
        parcela_da_tv:=valor_da_tv/parcelas
        fmt.Printf("%.2f\n%.2f\n", parcela_da_tv, valor_da_tv)
    } else {
        fmt.Printf("%.2f\n%.2f\n", valor_da_tv, valor_da_tv)
    }
}

