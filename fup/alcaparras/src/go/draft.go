package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
    l:= bufio.NewReader(os.Stdin)
    texto, _:= l.ReadString('\n')
    letraStr, _:= l.ReadString('\n')
    
    letra:= []rune(letraStr)[0]
    cont:=0
    for _, c:= range texto{
        if c== letra{
            cont++
        }
    }
    fmt.Println(cont)
}