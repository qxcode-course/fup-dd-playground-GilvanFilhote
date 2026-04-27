package main
import "fmt"
import "math"
func main() {
    var c, a float64
    fmt.Scan(&c, &a)
    var viagens float64
    viagens=a/(c-1)
    fmt.Printf("%d\n", int(math.Ceil(viagens)))

}
