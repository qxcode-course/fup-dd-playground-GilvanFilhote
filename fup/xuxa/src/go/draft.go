package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    texto, _:= leitor.ReadString('\n')
    texto= strings.TrimSpace(texto)
    letras:= []rune(texto)
    for i, j :=0, len(letras)-1; i<j; i, j= i+1, j-1{
        letras[i], letras [j]=letras[j], letras[i]
        
    }
    fmt.Println(string(letras))

}