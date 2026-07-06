package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    frase = strings.TrimSpace(frase)
    palavras:= strings.Split(frase, " ")
    for i:= 0; i<len(palavras)-1; i++{
        if palavras[i]>palavras[i+1]{
            fmt.Println("nao")
            return
        }
    }





    fmt.Println("sim")
}