package main
import "fmt"
import "math"
func main() {
    var tempo, hora, minuto, segundo float64
    fmt.Scan(&tempo)
    hora = int(tempo/3600)
    restodehoras := math.Mod (tempo, 3600)
    minuto= int (restodehoras/60)
    segundo= math.Mod(tempo, 60)
    fmt.Printf("%d:%d:%.2d", hora, minuto, segundo)
}
