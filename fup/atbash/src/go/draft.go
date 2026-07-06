package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
    l:= bufio.NewReader (os.Stdin)
    texto, _:= l.ReadString('\n')
    p1, _:= l.ReadString('\n')
    p2, _:= l.ReadString('\n')
    texto= texto[:len(texto)-1]
    p1= p1[:len(p1)-1]
    p2= p2[:len(p2)-1]
    m:= make(map[rune]rune)
    for i:=0; i<len(p1); i++{
        m[rune(p1[i])]= rune(p2[i])
        m[rune(p2[i])]= rune(p1[i])
    }
    for _, c:= range texto{
        if troca, existe:= m[c]; existe{
            fmt.Print(string(troca))
        } else{
            fmt.Print(string(c))
        }
    }
    fmt.Println()
}