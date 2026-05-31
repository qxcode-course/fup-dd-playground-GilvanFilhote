package main
import "fmt"
func main() {
    var qtd, choco, limao, tarde, manha int
    fmt.Scan(&qtd)
    vet_sabor:=make([]string, qtd)
    vet_turno:=make([]string, qtd)

    for i:=0; i<qtd; i++{
        fmt.Scan(&vet_sabor[i])
        fmt.Scan(&vet_turno[i])
        if vet_sabor[i]=="c"{
            choco+=1
        } else if vet_sabor[i]=="l"{
            limao+=1
        }
        if vet_turno[i]=="m"{
            manha+=1
        } else if vet_turno[i]=="t"{
            tarde+=1
        } 
    }
        if choco>limao{
            fmt.Println("c")
        } else if choco<limao{
            fmt.Println("l")
        } else {
            fmt.Println("empate")
        }
        if manha>tarde{
            fmt.Println("t")
        } else if tarde>manha{
            fmt.Println("m")
        } else {
            fmt.Println("empate")
        }


}