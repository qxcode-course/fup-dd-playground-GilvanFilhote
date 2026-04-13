package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Print("[ ")
    
    for i:=0; i<n; i++{
        fmt.Printf("%d ", i)

    
    }
    for i:=n+1; i<10; i++{
        fmt.Printf("%d ", i)
        if i==10{
            break
        }
    } 
    fmt.Println("ceu ]")
}

