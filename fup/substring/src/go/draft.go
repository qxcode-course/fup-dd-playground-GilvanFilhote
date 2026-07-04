package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
    
    leitor:= bufio.NewReader(os.Stdin)
    leitor.ReadString('\n')
    texto, _:= leitor.ReadString('\n')
    
    texto= strings.TrimSpace(texto)
    
    var indice, tam_substring int
    fmt.Scan(&indice, &tam_substring)

    fim:= indice+tam_substring
    if fim>len(texto){
        fim=len(texto)
    }

    
    fmt.Printf(texto[indice:fim])


    
}