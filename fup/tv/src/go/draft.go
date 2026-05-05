package main
import "fmt"
func main() {
    var valor, parcela_s float64
    fmt.Scan(&valor, &parcela_s)
    valor_total:=0
    valor_da_parcela:=0
    if parcela_s==1{
        parcela_s=parcela_s-1
        valor_da_parcela:=valor*parcela_s*0.05
        valor_total:=valor_da_parcela*parcela_s
        fmt.Printf("%f\n%f\n", valor_total, valor_da_parcela)
    } else {
        fmt.Printf("%f\n%f\n", valor_total, valor_da_parcela)
    }

}

