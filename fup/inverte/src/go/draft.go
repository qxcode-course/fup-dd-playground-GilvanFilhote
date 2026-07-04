
package main
import "fmt"
func main() {
    var letra rune
    fmt.Scanf("%c", &letra)
    if letra >='a' && letra <='z'{
        letra-=32
    } else if letra>='A' && letra<='Z'{
        letra+=32
    }
    fmt.Printf("%c\n", letra)
}