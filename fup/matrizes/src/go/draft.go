package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    mat1:= make ([][]int, a)
    mat2:= make ([][]int, a)
    soma:= make ([][]int, a)
    

    for i:=0; i<a; i++{
        mat1[i]= make ([]int, b)
        mat2[i]= make ([]int, b)
        soma[i]= make ([]int, b)
    }
    for i:=0; i<a; i++{
        for j:=0; j<b; j++{
            fmt.Scan(&mat1[i][j])
        }
    }
    for i:=0; i<a; i++{
        for j:=0; j<b; j++{
            fmt.Scan(&mat2[i][j])
        }
    }
    for i:=0; i<a; i++{
        for j:=0; j<b; j++{
            soma[i][j]=mat1[i][j]+mat2[i][j]
        }
    }

    for i:=0; i<a; i++{
        fmt.Print("[ ")
    for j:=0; j<b; j++{
         fmt.Print(soma[i][j], " ")   
        }
        fmt.Println("]")
    }
}

