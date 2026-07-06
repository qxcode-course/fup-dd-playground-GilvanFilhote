package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    lei:= bufio.NewReader(os.Stdin)
    frase, _:= lei.ReadString('\n')
    palavra:=strings.Fields(frase)
    for i, p:= range palavra{
        for j:=0; j<len(p)-1; j++{
        if strings.Contains("AEIOUaeiou", string(p[j])) && !strings.Contains("AEIOUaeiou", string(p[j+1])){
            s:= p[:j+1]
            palavra[i]= s + s + p
            break
        }
        
    }

    }


    fmt.Println(strings.Join (palavra, " "))
}