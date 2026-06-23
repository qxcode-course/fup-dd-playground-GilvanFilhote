package main
import "fmt"
func ehprimo(x int) int{
    resul:=1
    for i:=2; i<x; i++{
        if x%i==0{
            resul=0
        
        }
    }
    return resul;
}
func main() {
    var num int
    fmt.Scan(&num)
    resultado:= ehprimo(num)
    fmt.Println(resultado)
}