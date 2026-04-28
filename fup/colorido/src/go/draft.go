package main
import "fmt"
func main() {
    var n int
    var l string
    fmt.Scan(&n, &l)
    fmt.Printf("[ ")
    for i := 0; i <= 10; i++ {
        if l == "d" {
    if i%2==0 || i==0 {
        fmt.Printf("%dd ", i)
    }    
} else if l == "e" {
    if i%2==0 || i==0 {
        fmt.Printf("%de ", i)
    }    
}

}
}
