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
    codigo, _:= leitor.ReadString('\n')
    subst, _:= leitor.ReadString('\n')
    verdadeira:= strings.Replaceall

    fmt.Println("Hello, World!")
}