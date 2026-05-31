package main
import "fmt"

func main() {
    
   // 𝑁: Um número inteiro que indica a dimensão do tabuleiro.
//𝑋: Um número inteiro representando a posição inicial horizontal da cabeça da cobra.
//𝑌: Um número inteiro representando a posição inicial vertical da cabeça da cobra.
//𝐶: Um caractere representando a direção da cobra ('U' para cima, 'D' para baixo, 'L' para esquerda, 'R' para direita).
//𝑆: Um número inteiro representando o número de segundos de distração.
var n, x, y, s int
var c string
fmt.Scan(&n, &x, &y, &c, &s)

switch c{
case "U":
    y=y-s
case "D":
    y=y+s
case "L":
    x=x-s
case "R":
    x=x+s
}
x = ((x % n) + n) % n
y = ((y % n) + n) % n
fmt.Printf("%d %d\n", x, y)
}
