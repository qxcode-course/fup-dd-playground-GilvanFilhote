package main
import "fmt"
func main() {
    matriz:=[2][3]int{}
    for i:=0; i<2; i++{
        for j:=0; j<3; j++{
            fmt.Scan(&matriz[i][j])
        }
    }
    soma:=0
    for i:=0; i<2; i++{
        for j:=0; j<3; j++{
           soma+= matriz[i][j]
        }
    }
    fmt.Println(soma)
}