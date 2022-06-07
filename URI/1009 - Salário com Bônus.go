package main
 
import (
    "fmt"
)
 
func main() {
    var sal_fixo, vendas float64
    var nome string
    
    fmt.Scanln(&nome)
    fmt.Scanln(&sal_fixo)
    fmt.Scanln(&vendas)
    
    fmt.Printf("TOTAL = R$ %.2f\n", sal_fixo + vendas * 15 / 100)
}
