package main
import "fmt"
func main(){
    var A, B, C, M int
    fmt.Scan(&M, &A, &B)
    C = M - (A+B)
    if A > B && A>C {
        fmt.Printf("%d\n", A)
    } else if B > A && B>C {
        fmt.Printf("%d\n", B)
    } else if C > B && C > A {
        fmt.Printf("%d\n", C)
    }







}
    