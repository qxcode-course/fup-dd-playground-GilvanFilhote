package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"

)
func main() {
    l:= bufio.NewReader(os.Stdin)
    frase, _:= l.ReadString('\n')
    palavras:= strings.Fields(frase)
    soma:=0
    for _, p := range palavras{
        if num, err:= strconv.Atoi(p); err==nil{
            soma+=num
        }
    }
    fmt.Println(soma)

}