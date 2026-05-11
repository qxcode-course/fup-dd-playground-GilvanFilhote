package main
import "fmt"
func main() {
    var t int
    fmt.Scan(&t)
    if t==0{
        fmt.Println("[]")
    } else{
    arr:= make([]int, t)
    for i:=0; i<t; i++{
        fmt.Scan(&arr[i])
    }
    fmt.Printf("[")
    for i:=0; i<t; i++{
        if i!=0{
        fmt.Printf(", ")
        }

        fmt.Printf("%d", arr[i])
    }
    fmt.Println("]")
}
}
