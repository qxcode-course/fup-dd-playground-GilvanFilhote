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

    codigo, _:= leitor.ReadString('\n')
    codigo= strings.TrimSpace(codigo)
    subst, _:= leitor.ReadString('\n')
    subst= strings.TrimSpace(subst)

    verdadeira:= strings.ReplaceAll(frase, codigo, subst)

    fmt.Println(verdadeira)
}