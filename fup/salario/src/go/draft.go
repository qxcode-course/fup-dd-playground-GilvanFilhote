package main
import "fmt"
func main() {
    var salario float64
    fmt.Scan(&salario)
    if salario<=1000 {
        novosalario:=salario+(salario*20/100)
        fmt.Printf("%.2f\n", novosalario)
    } else if salario>1000 && salario<=1500 {
        novosalario:=salario+(salario*15/100)
        fmt.Printf("%.2f\n", novosalario)
    } else if salario>1500 && salario<=2000 {
        novosalario:=salario+(salario*10/100)
        fmt.Printf("%.2f\n", novosalario)
    } else if salario>2000 {
        novosalario:=salario+(salario*5/100)
        fmt.Printf("%.2f\n", novosalario)
    }

}
