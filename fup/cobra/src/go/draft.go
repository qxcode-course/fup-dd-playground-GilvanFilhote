package main
import "fmt"
import "math"
func main() {
    
   // 𝑁: Um número inteiro que indica a dimensão do tabuleiro.
//𝑋: Um número inteiro representando a posição inicial horizontal da cabeça da cobra.
//𝑌: Um número inteiro representando a posição inicial vertical da cabeça da cobra.
//𝐶: Um caractere representando a direção da cobra ('U' para cima, 'D' para baixo, 'L' para esquerda, 'R' para direita).
//𝑆: Um número inteiro representando o número de segundos de distração.
var n, x, y, s float64
var c string
fmt.Scan(&n, &x, &y, &c, &s)
switch c{
case "u"{
    y=y-s
}
case "d"{
    y=y+s
}
case "l"{
    x=x-s
}
case "r"{
    x=x+s
}
}
fmt.Printf("%.0f %.0f\n", x, y)
}
