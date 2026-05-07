package main
import "fmt"
func main() {
   var h, p, f, d int
   fmt.Scan(&h, &p, &f, &d)
   i:=f
   for {
    switch d {
    case 1:
      i++
      if i==16{
        i=0
      }
      
       case -1:
        i--
        if i==-1{
          i=15
        }
    }
      
      if i==h{
        fmt.Println("S")
        break
        } else if i==p{
          fmt.Println("N")
          break
        }
      }
}