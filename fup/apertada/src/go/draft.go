package main
import "fmt"
func main() {
        vetor:=make([]int, 5)
        menorvalor:=0
        for i:=0; i<5; i++{
            fmt.Scan(&vetor[i])
            if menorvalor>vetor[i]||i==0{
                menorvalor=vetor[i]
            }
        }
        fmt.Printf("%d\n", menorvalor)
}
