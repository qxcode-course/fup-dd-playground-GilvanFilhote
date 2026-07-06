package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "unicode"
)
func main() {
    l:=bufio.NewReader(os.Stdin)
    f, _:=l.ReadString('\n')
    c, _:=l.ReadString('\n')
    m, _:=l.ReadString('\n')
    r:= []rune(m)
    for _, x:= range strings.TrimSpace(f){
        if !unicode.IsLetter(x)||strings.ContainsRune(c, unicode.ToLower(x)){
            fmt.Print(string(x))
        } else{
            fmt.Print(string(r[0]))
        }
    }

    fmt.Println("")
}