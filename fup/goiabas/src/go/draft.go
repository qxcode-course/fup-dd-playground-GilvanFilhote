package main
import "fmt"
func main() {
   var a, b, c, d int
   fmt.Scan(&a, &b, &c, &d)
   tempo:=(b+c+d+a-1)/a
   fmt.Printf("%d\n", tempo)
}
