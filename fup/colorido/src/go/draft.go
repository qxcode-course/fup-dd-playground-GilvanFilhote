package main
import "fmt"
func main() {
    var n int
    var l string
    fmt.Scan(&n, &l)
    fmt.Printf("[ ")

    for i:=0; i<=10; i++ {
        par:=(i%2==0)
        sufixo:=""
        if i==n || i==10{
            continue
        }

        if l =="d"{
            if par {
                sufixo= "d"
            } else{
                sufixo= "e"
            } 
        } else {
            if par {
                sufixo= "e"
            } else{
                sufixo= "d"
            } 
        }

        if i>n {
            if sufixo=="d"{
                sufixo="e"
            } else {
              sufixo="d"  
            }
        }
        
        fmt.Printf("%d%s ", i, sufixo)
    }
    fmt.Print("ceu ]\n")
}