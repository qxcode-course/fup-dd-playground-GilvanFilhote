package main
import "fmt"
func main() {
   var h, p, f, d int
   fmt.Scan(&h, &p, &f, &d)
  if d==-1{
    for i:=f; i<=15; i++{
      if i>15{
        i=0
      } else if i==h{
        fmt.Println("S")
        break
      } else if i==p{
        fmt.Println("N")
        break
      }
    }
  } else if d==1{
    for i:=f; i>=0; i--{
      if i<0{
        i=15
      } else if i==h{
        fmt.Println("S")
        break
      } else if i==p{
        fmt.Println("N")
        break
      }
    }
  }
  }