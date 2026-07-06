package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)
func main() {
    in:=bufio.NewScanner(os.Stdin)

    in.Scan()
    texto:= in.Text()

    in.Scan()
    indice, _:= strconv.Atoi(in.Text())

    in.Scan()
    tam, _ := strconv.Atoi(in.Text())

    fim:= tam+indice
    if indice>=len(texto){
        return
    }
    if indice+tam>=len(texto){
        fim= len(texto)
    }


    
    fmt.Printf(texto[indice:fim])


    fmt.Println("")
}