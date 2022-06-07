package main
import "fmt"

func main() {
    pi := 3.14159
    var r float64
    
    fmt.Scanln(&r)
    a := pi*r*r
    fmt.Printf("A=%.4f\n", a)
}
