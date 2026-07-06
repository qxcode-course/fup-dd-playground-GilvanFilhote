package main
import "fmt"
func main() {
    var palavra1, palavra2 string
    fmt.Scan(&palavra1, &palavra2)
    n:=len(palavra1)
    if n>len(palavra2){
        n=len(palavra2)
    }
    k:=0
    for i:=1; i<n; i++{
        if palavra1[len(palavra1)-i:]==palavra2[:i]{
            k=i
        }
    }
    fmt.Print(palavra1[:len(palavra1)-k]+palavra2[k:])
}