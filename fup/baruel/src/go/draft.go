package main
import "fmt"
func main() {
    var quant_album, quant_fig_baruel int
    fmt.Scan(&quant_album, &quant_fig_baruel)
    cont:= make([]int, quant_album+1)

    for i:=0; i<quant_fig_baruel; i++{
        var fig int
        fmt.Scan(&fig)
        cont[fig]++
    }
    
    fmt.Printf("[")
    for i:=1; i<=quant_album; i++{
        if cont[i]>1{
        for j:=0; j<cont[i]-1; j++{
        fmt.Print(" ", i)
        
        }
    }
}
    fmt.Printf(" ]\n")
    fmt.Printf("[")
    
    for i:=1; i<=quant_album; i++{
        if cont[i]==0{
        fmt.Print(" ", i)
       
        }

    }
    fmt.Printf(" ]\n")

}