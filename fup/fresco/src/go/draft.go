package main
import (
    "fmt"
    "os"
    "bufio"


)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    
    fmt.Println("Hello, World!")
}