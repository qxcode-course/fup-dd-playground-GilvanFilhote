package main
import "fmt"
func main() {
    var valor, parcela_s float64
    fmt.Scan(&valor, &parcela_s)
    if valor>1{
        parcela_s=parcela_s-1
        valordaparcela:=valor*parcela_s*0.05
    } else {
        valordaparcela:=valor*parcela_s*0.05
    }

    
    fmt.Printf("%f\n", valordaparcela)

}

