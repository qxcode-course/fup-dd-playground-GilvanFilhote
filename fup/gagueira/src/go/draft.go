package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    frase= strings.TrimSpace(frase)
    palavras:= strings.Fields(frase)
    for i, p:= range palavras{
        if i>0{
            fmt.Print(" ")
        }
        fmt.Print(p, " ", p)
    }
    fmt.Println("")
}