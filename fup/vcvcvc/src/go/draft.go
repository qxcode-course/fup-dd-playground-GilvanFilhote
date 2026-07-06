package main
import (
    "fmt"
    "os"
    "bufio"

)
func main() {
        leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    runes:= []rune(frase)

    for i:=0; i<len(runes); i++{
        
        if runes[i]=='A' || runes[i]=='E' ||runes[i]=='I' ||runes[i]=='O' ||runes[i]=='U' ||runes[i]=='a' ||runes[i]=='e' ||runes[i]=='i' ||runes[i]=='o' ||runes[i]=='u'{
            runes[i]='v'
        } else if (runes[i]>='b' && runes[i]<='z') || (runes[i]>='B' && runes[i]<='Z'){
            runes[i]='c'
        }

    }
    frase= string(runes)
    fmt.Printf(frase)

}