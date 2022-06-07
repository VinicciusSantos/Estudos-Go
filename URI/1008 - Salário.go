package main
 
import (
    "fmt"
)
 
func main() {
    var num, horas, sal_h float64
    
    fmt.Scanln(&num)
    fmt.Scanln(&horas)
    fmt.Scanln(&sal_h)
    
    sal_tot := horas * sal_h
    fmt.Printf("NUMBER = %v\n", num)
    fmt.Printf("SALARY = U$ %.2f\n", sal_tot)
}
