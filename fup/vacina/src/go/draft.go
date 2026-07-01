package main
import (
    "fmt"
    "slices"
)

func main() {
   var vacinas int
   fmt.Scan(&vacinas)
   força:= make([]int, vacinas)
   celulas:= make([]int, vacinas)
   for i:=0; i<vacinas; i++{
    fmt.Scan(&força[i])
    // fmt.Scan(&celulas[i])
   }
   for i:=0; i<vacinas; i++{
   
     fmt.Scan(&celulas[i])
   }
   slices.Sort(força)
   slices.Sort(celulas)
   salvatudo:=true
   for i:=0; i<vacinas; i++{
    if força[i]<celulas[i]{
        salvatudo=false
    }
}
if salvatudo==true{
    fmt.Printf("Yes\n")
} else{
    fmt.Printf("No\n")
}
   
}