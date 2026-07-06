package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    matriz:= make([][]int, n)
    for i:=0; i<n; i++{
        matriz[i]= make([]int, n)
    }
    maior:=-1
    colmaior:=0
    for i:=0; i<n; i++{
        for j:=0; j<n; j++{
            fmt.Scan(&matriz[i][j])
        }   
    }
    for i:=0; i<n; i++{
        soma:=0
        for j:=0; j<n; j++{
            soma+= matriz[j][i]*matriz[j][i] 
        }
        if soma>maior{
            maior=soma
            colmaior=i
        }
    }
    for j:=0; j<n; j++{
        matriz[j][colmaior]*=matriz[j][colmaior]
    }
    fmt.Println(colmaior)
}