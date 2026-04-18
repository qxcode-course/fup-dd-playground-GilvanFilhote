package main
import "fmt"
func main() {
    var tempo int
    var restodehoras, hora, minuto, segundo int
    fmt.Scan(&tempo)
    hora = tempo/3600
    restodehoras = tempo%3600
    minuto = restodehoras/60
    segundo= restodehoras%60
    fmt.Printf("%d:%d:%d\n", hora, minuto, segundo)
}
