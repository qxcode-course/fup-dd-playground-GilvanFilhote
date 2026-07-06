package main
import "fmt"
func main() {
    vet:= make([]int, 6)
    for i:=0; i<6; i++{
        fmt.Scan(&vet[i])
    }   
    matriz:=[4][4]int{
        {1, 9, 27, 23},
        {34, 20, 37, 47},
        {30, 87, 55, 69},
        {13, 60, 99, 66},
    }
    cont:=0
    for i:=0; i<=3; i++{
        for j:=0; j<=3; j++{
            for k:=0; k<6; k++{
                if matriz[i][j]==vet[k]{
                    cont+=1
                    break
                }

            }
        }
    }
    fmt.Println(cont)
}