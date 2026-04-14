package main
import "fmt"
func main() {
    var quantidade int
    fmt.Scan(&quantidade)
    var arr[]int= make([]int, quantidade)
    if quantidade == 0{
        fmt.Print("\n")
    }
    for quantidade:= range arr {
        fmt.Scan(&arr[quantidade])
    }
    for _, valores:= range arr {
        fmt.Printf("%d\n", valores)
    }

   
   
}
