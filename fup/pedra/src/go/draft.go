package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    vet:= make([][2]int, n)

    for i:=0; i<n; i++{
        for j:=0; j<2; j++{
            fmt.Scan(&vet[i][j])
        }
    }
    
    indice:=0
    for i:=0; i<n; i++{
        for j:=0; j<2; j++{
            if vet[i][j]>=10{
                if (vet[i][j]+vet[i][j+1])<(vet[i+1][j]+vet[i+1][j+1]){
                    // vence:=(vet[i][j]+vet[i][j+1])
                    indice=i
                }
            } else {
                fmt.Println("sem ganhador")
            }
            
        }
    }
    fmt.Print(indice)
}