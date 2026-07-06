package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    )

func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    subfrase, _:= leitor.ReadString('\n')

    frase=strings.TrimSpace(frase)
    subfrase=strings.TrimSpace(subfrase)
    cont:= strings.Count(frase, subfrase)
    fmt.Printf("%d\n", cont)
}