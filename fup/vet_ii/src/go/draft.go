package main
import "fmt"
func main() {
    var quantidade int
    fmt.Scan(&quantidade)
    vetorr := make ([]int, quantidade)
    for _, valor := range vetorr {
    fmt.Scan(&vetorr[valor])
}
    fmt.Printf("[ ")
    for _, valor := range vetorr {
    fmt.Printf("%d ", valor)
}
    fmt.Printf("]\n")
}
