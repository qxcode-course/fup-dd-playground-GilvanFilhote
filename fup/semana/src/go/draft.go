package main
import "fmt"
func main() {
    var dia, horario int
    fmt.Scan(&dia, &horario)
    if (dia >=2 && dia<=6) && ((horario>=8 && horario<=11) ||( horario>=14 && horario<=17)){
        fmt.Println("SIM")
    } else if (dia == 7) && (horario>=8 && horario<=11){
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
