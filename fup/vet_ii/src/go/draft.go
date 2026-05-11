package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
        if qtd==0{
        fmt.Println("[ ]")
        return
    }
        fmt.Printf("[ ")
        var arr []int = make ([]int, qtd)
        for i:=0 ; i<qtd; i++{
        fmt.Scan(&arr[i])
    }
    
        for i:= 0; i<qtd; i++{
            fmt.Printf("%d ", arr[i])
        }
        fmt.Println("]")
}
