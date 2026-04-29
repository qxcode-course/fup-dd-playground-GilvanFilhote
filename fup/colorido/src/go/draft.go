package main
import "fmt"
func main() {
    var n int
    var l string
    fmt.Scan(&n, &l)
    fmt.Printf("[ ")

    if l=="d"{
    for i := 0; i <= 10; i++ {
       if i==n{
        continue
       } 
       if i==10{
        fmt.Print("ceu")
        if i<n{
        fmt.Printf("%dd ", i)
        }else if i%2==0 {
        fmt.Printf("%de ", i)
       }
       }else{
        fmt.Printf("%de ", i)
       }
    }
}
fmt.Printf("]")
}