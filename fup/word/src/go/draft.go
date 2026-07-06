package main
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "unicode"
    )
func main() {
    l:= bufio.NewReader(os.Stdin)
    text, _:= l.ReadString('\n')
    text= strings.TrimSpace(text)
    op, _:= l.ReadString('\n')
    op= strings.TrimSpace(op)
   
    switch op{
    case "M":
        fmt.Print(strings.ToUpper(text))
    
    case "m":
        fmt.Print(strings.ToLower(text))

    case "i":
        for _, x:= range text{
            if unicode.IsUpper(x){
                fmt.Print(string(unicode.ToLower(x)))
            } else if unicode.IsLower(x){
                fmt.Print(string(unicode.ToUpper(x)))
            } else{
                fmt.Print(string(x))
            }
        }


    case "p":
        p:=strings.Fields(strings.ToLower(text))
        for i, x:= range p{
            if len(x)>1{
                x=strings.ToUpper(x[:1])+x[1:]
            }
            if i>0{
                fmt.Print(" ")
            }
            fmt.Print(x)
        }
    }
    fmt.Println()
}
