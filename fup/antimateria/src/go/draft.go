package main
import (
    "fmt"
    "strings"
    
)
func main() {
    var palavra1, palavra2 string
    fmt.Scan(&palavra1, &palavra2)

    n:=len(palavra1)

    if n>len(palavra2){
        n=len(palavra2)
    }
    k:=0
    pos:=0
    for i:=1; i<=n; i++{
        s:= palavra1[len(palavra1)-i:]
        p:=strings.Index(palavra2, s)
    
        if p!=-1{
            k=i
            pos=p
        }
    }
    
    fmt.Print(
        palavra1[:len(palavra1)-k]+
        palavra2[:pos]+
        palavra2[pos+k:],
    )
    fmt.Println()
}