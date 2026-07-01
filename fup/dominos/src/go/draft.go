package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    domino:= make([]int, n)
    for i:=0; i<n; i++{
        fmt.Scan(&domino[i])
    }
    ordenado:=true
    for i:=0; i<n-1; i++{
        if domino[i]>domino[i+1]{
            ordenado=false
        }
    }
    if ordenado==true{
        fmt.Println("ok")
    } else{
        fmt.Println("precisa de ajuste")
    }
}