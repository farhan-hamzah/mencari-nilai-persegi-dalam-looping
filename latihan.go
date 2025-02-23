package main
import "fmt"

func main(){
    var x, y, z, hasil, hasil2 int
    fmt.Scan(&x)
    z = 0
    for z < x {
        fmt.Scan(&y)
        hasil2 = y*y
        hasil = x*y
        fmt.Print(hasil, hasil2)
        z++      
    }
}