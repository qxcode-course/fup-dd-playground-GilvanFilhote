package main
import (
  "fmt"
 "strconv"
)

func main() {
    var num, num1 int
    fmt.Scan(&num, &num1)
    cont:=0
    cacado:=strconv.Itoa(num)
    aparecido:=strconv.Itoa(num1)
    for i:=0; i<len(aparecido); i++{
        if cacado== string(aparecido[i]){
            cont++
        }
    }
    fmt.Printf("%d\n", cont)
    
    }
