package main
import "fmt"
func main() {
    matriz:= [3][3]int{}
    for i:=0; i<3; i++{
        for j:=0; j<3; j++{
        fmt.Scan(&matriz[i][j])

    }
}
alvo:= matriz[0][0]+matriz[0][1]+matriz[0][2]
d1, d2:=0, 0
magico:=true
    for i:=0; i<3; i++{
        linha, coluna:= 0, 0
        for j:=0; j<3; j++{
            linha+= matriz[i][j]
            coluna+= matriz[j][i]
        }
        if linha!=alvo||coluna!=alvo{
            magico=false
        }
        d1+=matriz[i][i]
        d2+=matriz[i][2-i]
    }
    if d1!=alvo||d2!=alvo{
        magico=false
    }
    if magico==false{
    fmt.Println("nao")
    } else{
        fmt.Println("sim")
    }

}