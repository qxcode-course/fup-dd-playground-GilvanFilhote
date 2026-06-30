package main
import "fmt"
func main() {
    var v int
    fmt.Scan(&v)
    vet:= make([]int, v)
    for i:=0; i<v; i++{
        fmt.Scan(&vet[i])
    }
    sozero:=true
    for i:=0; i<v; i++{
        if vet[i]!=0{
            sozero=false
        }
    }
    if sozero{
        fmt.Printf("0\n")
    } else{
    for i:=0; i<v; i++{
       
        fmt.Printf("%d", vet[i])
        
        }
        fmt.Printf("\n")
    }
}
