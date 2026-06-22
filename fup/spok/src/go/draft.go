package main
import (
    "fmt"
    "strconv"
)
func main() {
    var numero int
    fmt.Scan(&numero)
    palindromo:=strconv.Itoa(numero)
    ehpalindromo:=1
    for i:=0; i<len(palindromo)/2; i++{
        if palindromo[i]!=palindromo[len(palindromo)-1-i]{
            ehpalindromo=0
        }
    }
    if ehpalindromo==1{
        fmt.Println(1)
    } else if ehpalindromo==0{
        fmt.Println(0)
    }
        
    }
