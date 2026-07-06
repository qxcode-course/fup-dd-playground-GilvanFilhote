package main
import (
    "os"
    "fmt"
    "bufio"

)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    runes:= []rune(frase)

    for i:=0; i<len(runes); i++{
        if (runes[i]>='a' && runes[i]<='z'){
            runes[i]-=32
        } else if (runes[i]>='A' && runes[i]<='Z'){
            runes[i]+=32
        }
    }
    frase= string(runes)
    fmt.Printf(frase)
}