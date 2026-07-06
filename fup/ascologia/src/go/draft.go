package main
import "fmt"
func main() {
    var nome string
    fmt.Scan(&nome)
    soma:= 0
    for _, c:= range nome{
        if c!= '\n' {
            soma+= int(c)
        }
    }
    fmt.Println(soma%50)
}