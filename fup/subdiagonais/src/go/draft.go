package main
import "fmt"
func main() {
    matriz:= [5][5]int{}
    somadiag:=0
    for i:=0; i<5; i++{
        for j:=0; j<5; j++{
            fmt.Scan(&matriz[i][j])
        }
    }
    somadiag+= (matriz[0][0]+matriz[1][1]+matriz[2][2]+matriz[3][3]+matriz[4][4])-(matriz[0][4]+matriz[1][3]+matriz[2][2]+matriz[3][1]+matriz[4][0])
    fmt.Println(somadiag)
}