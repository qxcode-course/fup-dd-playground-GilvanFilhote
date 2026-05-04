package main
import "fmt"
func main() {
    var velo_media, tempo_minutos, consumo_de_gaso int
    fmt.Scan(&velo_media, &tempo_minutos, &consumo_de_gaso)
    tempo_horas:= tempo_minutos/60
    distancia:= tempo_horas*velo_media
    desempenhofinal:= distancia/consumo_de_gaso
    fmt.Printf("%.2d", desempenhofinal)
}
