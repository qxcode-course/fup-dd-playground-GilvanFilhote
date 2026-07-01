package main
import "fmt"
func main() {
   var v int
   fmt.Scan(&v)
   vetor:= make([]int, v)
   for i:=0; i<v; i++{
    fmt.Scan(&vetor[i])
}
fmt.Printf("[")
    for i:=0; i<v; i++{
        if i!=v && i!=0 {
            fmt.Printf(", ")
        }

        if vetor[i]==1{
            fmt.Printf("A")
        } else if vetor[i]==11{
            fmt.Printf("J")
        } else if vetor[i]==12{
            fmt.Printf("Q")
        } else if vetor[i]==13{
            fmt.Printf("K")
        } else {
            fmt.Printf("%d", vetor[i])
        }
    }
    fmt.Printf("]\n")

}