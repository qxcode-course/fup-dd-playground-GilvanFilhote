package main
import "fmt"
func main() {
    var v int
    fmt.Scan(&v)
    vetor:= make([]float64, v)
    var media float64
    for i:=0; i<v; i++{
        fmt.Scan(&vetor[i])
        media+=vetor[i]
    }
    v1:= float64(v)
    media=media/v1
    fmt.Printf("%.2f\n", media)
    for i:=0; i<v; i++{
       
        if vetor[i]==media{
            fmt.Print("M")
        } else if vetor[i]>media {
            fmt.Printf("G")
        } else {
            fmt.Printf("P")
        }
        if i!=v-1{
            fmt.Printf(" ")
        }
    
    }
    fmt.Printf("\n")
}