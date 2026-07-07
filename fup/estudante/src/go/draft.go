package main
import (
    "fmt"
    "os"
    "bufio"
    "sort"
    "strings"

)

type Aluno struct{
    nome string
    nota1 float64
    nota2 float64
    nota3 float64
    media float64
}
func main() {
   
    l:= bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(l, &n)
    l.ReadString('\n')

    alunos:= make([]Aluno, n)

    for i:=0; i<n; i++{
    nome, _:= l.ReadString('\n')
    alunos[i].nome= strings.TrimSpace(nome)
    fmt.Fscan(l, &alunos[i].nota1, &alunos[i].nota2, &alunos[i].nota3)
    l.ReadString('\n')
    alunos[i].media= (alunos[i].nota1+alunos[i].nota2+alunos[i].nota3)/3
    }
    sort.Slice(alunos, func(i, j int)bool{
        return alunos[i].media>alunos[j].media

    })
    for i:= range alunos{
        fmt.Printf("%d: %s\n", i, alunos[i].nome)
        fmt.Printf("   Media: %.2f\n", alunos[i].media)
        fmt.Printf("   N1: %.2f, N2: %.2f, N3: %.2f\n", alunos[i].nota1, alunos[i].nota2, alunos[i].nota3)
    }

}