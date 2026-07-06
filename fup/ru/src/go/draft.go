package main
import 
(
"bufio"
"os"
"fmt"
)
func main() {
    leitor:= bufio.NewReader(os.Stdin)
    frase, _:= leitor.ReadString('\n')
    var vogais[] rune
    var conso[] rune

    for i:=0; i<len(frase); i++{
        letra:= rune(frase[i])
        if frase[i]=='a'||frase[i]=='e'||frase[i]=='i'||frase[i]=='o'||frase[i]=='u'||frase[i]=='A'||frase[i]=='E'||frase[i]=='I'||frase[i]=='O'||frase[i]=='U'{
            
            vogais= append(vogais, letra)
        } else if frase[i]>='a'&&frase[i]<='z' ||frase[i]>='A'&&frase[i]<='Z'{
            
            conso= append(conso, letra)
        }
    }
    fmt.Println(string(vogais))
    fmt.Println(string(conso))

}