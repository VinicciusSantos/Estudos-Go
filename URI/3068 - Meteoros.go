package main
 
import (
    "fmt"
)
 
func main() {
    var meteoros []int 
    var x1, y1, x2, y2, n int
    qtd_fazendas := 0
    
    for {
        quant_meteoritos := 0
        fmt.Scanln(&x1, &y1, &x2, &y2)
        if x1 == 0 && x2 == 0 && y1 == 0 && y2 == 0{
            break
        }
        
        fmt.Scanln(&n)
        
        for i := 0; i < n; i++ {
            var x, y int
            fmt.Scanln(&x, &y)
            if x >= x1 && x <= x2 && y <= y1 && y >= y2 {
                quant_meteoritos++
            }
        }
        
        meteoros = append(meteoros, quant_meteoritos)
        qtd_fazendas++
    }
    
    for i, fazenda := range meteoros{
        fmt.Printf("Teste %v\n", i+1)
        fmt.Println(fazenda)
    }
}
