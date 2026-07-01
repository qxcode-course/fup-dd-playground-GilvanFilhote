package main
import "fmt"
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)
    mmc:= num1
    for mmc%num2!=0{
        mmc+=num1
    }
    fmt.Println(mmc)
    }
